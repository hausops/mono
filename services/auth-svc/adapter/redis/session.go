package redis

import (
	"context"
	"errors"
	"fmt"
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

func (r *sessionRepository) DeleteByAccessToken(ctx context.Context, token session.AccessToken) error {
	sess, err := r.FindByAccessToken(ctx, token)
	if err != nil {
		return err
	}

	primaryKey := r.primaryKey(token)
	userIDKey := r.userIDKey(sess.UserID)
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		pipe.Del(ctx, primaryKey)
		pipe.Del(ctx, userIDKey)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey, userIDKey)
}

func (r *sessionRepository) FindByAccessToken(ctx context.Context, token session.AccessToken) (session.Session, error) {
	primaryKey := r.primaryKey(token)

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

		sess = session.Session{
			AccessToken: token,
			ExpireAt:    time.Unix(saved.ExpireAt, 0),
			UserID:      saved.UserID,
		}
		return nil
	}, primaryKey)

	return sess, err
}

func (r *sessionRepository) FindByUserID(ctx context.Context, userID string) (session.Session, error) {
	userIDKey := r.userIDKey(userID)

	var sess session.Session
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		accessTokenStr, err := tx.Get(ctx, userIDKey).Result()
		switch {
		case errors.Is(err, redis.Nil):
			return session.ErrNotFound
		case err != nil:
			return fmt.Errorf("get email from user ID %s: %w", userID, err)
		}

		token, err := session.ParseAccessToken(accessTokenStr)
		if err != nil {
			return fmt.Errorf("parse access token: %w", err)
		}

		sess, err = r.FindByAccessToken(ctx, token)
		if err != nil {
			return fmt.Errorf("FindByEmail(%s): %w", token, err)
		}
		return nil
	}, userIDKey)

	return sess, err
}

func (r *sessionRepository) Upsert(ctx context.Context, sess session.Session) error {
	primaryKey := r.primaryKey(sess.AccessToken)
	userIDKey := r.userIDKey(sess.UserID)
	// Watch the primary key and user ID key to detect changes by other clients.
	return r.client.Watch(ctx, func(tx *redis.Tx) error {

		// Get the current access token string for a given user ID
		// or an empty string.
		prevTokenStr, err := tx.Get(ctx, userIDKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.Get(%s): %w", userIDKey, err)
		}

		// Get the current user ID or empty string.
		prevUserID, err := tx.HGet(ctx, primaryKey, "userID").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, userID): %w", primaryKey, err)
		}

		pipe := tx.TxPipeline()

		// If a token already exists for the user, delete the old session
		// to ensure one active token per user.
		if prevTokenStr != "" {
			prevToken, err := session.ParseAccessToken(prevTokenStr)
			if err != nil {
				return fmt.Errorf("session.ParseAccessToken(prevToken: %s): %w", prevTokenStr, err)
			}
			pipe.Del(ctx, r.primaryKey(prevToken))
		}

		// If updating, remove the previous user ID index for the session.
		if prevUserID != "" {
			pipe.Del(ctx, r.userIDKey(prevUserID))
		}

		pipe.HSet(ctx, primaryKey, sessionRedis{
			ExpireAt: sess.ExpireAt.Unix(),
			UserID:   sess.UserID,
		})
		pipe.Set(ctx, userIDKey, sess.AccessToken.String(), 0)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey, userIDKey)
}

func (r *sessionRepository) primaryKey(token session.AccessToken) string {
	return fmt.Sprintf("auth-svc:session:%s", token)
}

func (r *sessionRepository) userIDKey(userID string) string {
	return fmt.Sprintf("auth-svc:session:user-id-idx:%s", userID)
}

// sessionRedis represents stored session data for a given key in redis.
type sessionRedis struct {
	// ExpireAt is stored as Unix timestamp e.g. 1687146872 (seconds) in redis.
	ExpireAt int64  `redis:"expireAt"`
	UserID   string `redis:"userID"`
}
