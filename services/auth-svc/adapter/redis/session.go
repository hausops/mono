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
	err = r.client.Watch(ctx, func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		pipe.Del(ctx, primaryKey)
		pipe.HDel(ctx, r.accessTokensKey(), sess.AccessToken.String())

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)

	if err != nil {
		return session.Session{}, err
	}
	return sess, nil
}

func (r *sessionRepository) FindByAccessToken(ctx context.Context, token session.AccessToken) (session.Session, error) {
	emailAddr, err := r.client.HGet(ctx, r.accessTokensKey(), token.String()).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return session.Session{}, session.ErrNotFound
	case err != nil:
		return session.Session{},
			fmt.Errorf("get email from access token %s: %w", token, err)
	}

	return r.FindByEmail(ctx, mail.Address{Address: emailAddr})
}

func (r *sessionRepository) FindByEmail(ctx context.Context, email mail.Address) (session.Session, error) {
	primaryKey := r.primaryKey(email)

	var saved sessionRedis
	err := r.client.HGetAll(ctx, primaryKey).Scan(&saved)
	switch {
	case errors.Is(err, redis.Nil):
		return session.Session{}, session.ErrNotFound
	case err != nil:
		return session.Session{}, fmt.Errorf("redis.HGetAll(%s): %w", primaryKey, err)
	}

	// Ensure that we check this condition after checking err != nil from HGetAll.
	// Otherwise, we'll always hit it when there's an error.
	if saved.AccessToken == "" {
		return session.Session{}, session.ErrNotFound
	}

	token, err := session.ParseAccessToken(saved.AccessToken)
	if err != nil {
		return session.Session{}, err
	}

	sess := session.Session{
		AccessToken: token,
		Email:       email,
		ExpireAt:    time.Unix(0, saved.ExpireAt),
	}
	return sess, nil
}

func (r *sessionRepository) Upsert(ctx context.Context, sess session.Session) error {
	primaryKey := r.primaryKey(sess.Email)
	// Watch the primary key to detect changes by other clients
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		prevToken, err := tx.HGet(ctx, primaryKey, "accessToken").Result()
		switch {
		case errors.Is(err, redis.Nil):
			// ok
		case err != nil:
			return fmt.Errorf("redis.HGet(%s, accessToken): %w", primaryKey, err)
		}

		pipe := tx.TxPipeline()
		pipe.HSet(ctx, primaryKey, sessionRedis{
			AccessToken: sess.AccessToken.String(),
			ExpireAt:    sess.ExpireAt.UnixNano(),
		})

		pipe.HDel(ctx, r.accessTokensKey(), prevToken)
		pipe.HSet(ctx, r.accessTokensKey(), sess.AccessToken.String(), sess.Email.Address)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)
}

func (r *sessionRepository) primaryKey(email mail.Address) string {
	return fmt.Sprintf("auth-svc:session-repo:email:%s", email.Address)
}

func (r *sessionRepository) accessTokensKey() string {
	return "auth-svc:session-repo:access-tokens"
}

// sessionRedis represents stored session data for a given key in redis.
type sessionRedis struct {
	AccessToken string `redis:"accessToken"`
	// ExpireAt is stored as UnixNano timestamp e.g. 1687146872522879000 in redis.
	ExpireAt int64 `redis:"expireAt"`
}
