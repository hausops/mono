package redis

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/hausops/mono/services/auth-svc/domain/session"
	"github.com/redis/go-redis/v9"
)

type sessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(c *redis.Client) *sessionRepository {
	return &sessionRepository{client: c}
}

var _ session.Repository = (*sessionRepository)(nil)

func (r *sessionRepository) DeleteByEmail(ctx context.Context, email mail.Address) (session.Session, error) {
	sess, err := r.FindByEmail(ctx, email)
	if err != nil {
		return session.Session{}, err
	}

	primaryKey := r.primaryKey(email)
	accessTokenKey := r.accessTokenKey(sess.AccessToken)
	err = r.client.Watch(ctx, func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		pipe.Del(ctx, primaryKey)
		pipe.Del(ctx, accessTokenKey)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey, accessTokenKey)

	if err != nil {
		return session.Session{}, err
	}
	return sess, nil
}

func (r *sessionRepository) FindByAccessToken(ctx context.Context, token session.AccessToken) (session.Session, error) {
	accessTokenKey := r.accessTokenKey(token)

	var sess session.Session
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		emailAddr, err := tx.Get(ctx, r.accessTokenKey(token)).Result()
		switch {
		case errors.Is(err, redis.Nil):
			return session.ErrNotFound
		case err != nil:
			return fmt.Errorf("get email from access token %s: %w", token, err)
		}

		email, err := mail.ParseAddress(emailAddr)
		if err != nil {
			return fmt.Errorf("parse email address: %w", err)
		}

		sess, err = r.FindByEmail(ctx, *email)
		if err != nil {
			return fmt.Errorf("FindByEmail(%s): %w", email.Address, err)
		}
		return nil
	}, accessTokenKey)

	return sess, err
}

func (r *sessionRepository) FindByEmail(ctx context.Context, email mail.Address) (session.Session, error) {
	primaryKey := r.primaryKey(email)

	var sess session.Session
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		// We need to check with Exists before calling HGetAll because,
		// when the key doesn't exist HGetAll returns nil error, instead of redis.Nil.
		// https://github.com/redis/go-redis/issues/1668
		n, err := tx.Exists(ctx, primaryKey).Result()
		if err != nil {
			return fmt.Errorf("redis.Exists(%s): %w", primaryKey, err)
		} else if n == 0 {
			return session.ErrNotFound
		}

		var saved sessionRedis
		err = tx.HGetAll(ctx, primaryKey).Scan(&saved)
		if err != nil {
			return fmt.Errorf("redis.HGetAll(%s): %w", primaryKey, err)
		}

		token, err := session.ParseAccessToken(saved.AccessToken)
		if err != nil {
			return err
		}

		sess = session.Session{
			AccessToken: token,
			Email:       email,
			ExpireAt:    time.Unix(0, saved.ExpireAt),
		}
		return nil
	}, primaryKey)

	return sess, err
}

func (r *sessionRepository) Upsert(ctx context.Context, sess session.Session) error {
	primaryKey := r.primaryKey(sess.Email)
	// Watch the primary key to detect changes by other clients
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		// Get the current access token or empty string
		prevTokenStr, err := tx.HGet(ctx, primaryKey, "accessToken").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, accessToken): %w", primaryKey, err)
		}

		pipe := tx.TxPipeline()
		pipe.HSet(ctx, primaryKey, sessionRedis{
			AccessToken: sess.AccessToken.String(),
			ExpireAt:    sess.ExpireAt.UnixNano(),
		})

		// If updating, remove the previous access token index for the session.
		if prevTokenStr != "" {
			prevToken, err := session.ParseAccessToken(prevTokenStr)
			if err != nil {
				return fmt.Errorf("parse previously stored access token: %w", err)
			}
			pipe.Del(ctx, r.accessTokenKey(prevToken))
		}

		pipe.Set(ctx, r.accessTokenKey(sess.AccessToken), sess.Email.Address, 0)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)
}

func (r *sessionRepository) primaryKey(email mail.Address) string {
	return fmt.Sprintf("auth-svc:session-repo:email:%s", email.Address)
}

func (r *sessionRepository) accessTokenKey(token session.AccessToken) string {
	return fmt.Sprintf("auth-svc:session-repo:access-token-idx:%s", token)
}

// sessionRedis represents stored session data for a given key in redis.
type sessionRedis struct {
	AccessToken string `redis:"accessToken"`
	// ExpireAt is stored as UnixNano timestamp e.g. 1687146872522879000 in redis.
	ExpireAt int64 `redis:"expireAt"`
}
