package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/redis/go-redis/v9"
)

type confirmRepository struct {
	client *redis.Client
}

func NewConfirmRepository(c *redis.Client) *confirmRepository {
	return &confirmRepository{client: c}
}

var _ confirm.Repository = (*confirmRepository)(nil)

func (r *confirmRepository) FindByToken(ctx context.Context, token confirm.Token) (confirm.Record, error) {
	tokenKey := r.tokenKey(token)

	var rec confirm.Record
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		userID, err := tx.Get(ctx, tokenKey).Result()
		switch {
		case errors.Is(err, redis.Nil):
			return confirm.ErrNotFound
		case err != nil:
			return fmt.Errorf("get email from token %s: %w", token, err)
		}

		rec, err = r.FindByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("FindByUserID(%s): %w", userID, err)
		}
		return nil
	}, tokenKey)

	return rec, err
}

func (r *confirmRepository) FindByUserID(ctx context.Context, userID string) (confirm.Record, error) {
	primaryKey := r.primaryKey(userID)

	var rec confirm.Record
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		confirmed, err := tx.HGet(ctx, primaryKey, "confirmed").Bool()
		switch {
		case errors.Is(err, redis.Nil):
			return confirm.ErrNotFound
		case err != nil:
			return fmt.Errorf("redis.HGet(%s, confirmed): %w", primaryKey, err)
		}

		tokenStr, err := tx.HGet(ctx, primaryKey, "token").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, token): %w", primaryKey, err)
		}

		var token confirm.Token
		if tokenStr != "" {
			token, err = confirm.ParseToken(tokenStr)
			if err != nil {
				return fmt.Errorf("confirm.ParseToken(%s): %w", tokenStr, err)
			}
		}

		rec = confirm.Record{
			IsConfirmed: confirmed,
			Token:       token,
			UserID:      userID,
		}

		return nil
	}, primaryKey)

	return rec, err
}

func (r *confirmRepository) Upsert(ctx context.Context, rec confirm.Record) error {
	primaryKey := r.primaryKey(rec.UserID)
	// Watch the primary key to detect changes by other clients
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		// Get the current token or empty string
		prevTokenStr, err := tx.HGet(ctx, primaryKey, "token").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, token): %w", primaryKey, err)
		}

		pipe := tx.Pipeline()

		// Remove the previous token index
		if prevTokenStr != "" {
			prevToken, err := confirm.ParseToken(prevTokenStr)
			if err != nil {
				return fmt.Errorf("parse previously stored token: %w", err)
			}
			pipe.Del(ctx, r.tokenKey(prevToken))
		}

		pipe.HSet(ctx, primaryKey, "confirmed", rec.IsConfirmed)

		if rec.Token.IsZero() {
			pipe.HDel(ctx, primaryKey, "token")
		} else {
			pipe.HSet(ctx, primaryKey, "token", rec.Token.String())
			pipe.Set(ctx, r.tokenKey(rec.Token), rec.UserID, 0)
		}

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)
}

func (r *confirmRepository) primaryKey(userID string) string {
	return fmt.Sprintf("auth-svc:confirm:%s", userID)
}

func (r *confirmRepository) tokenKey(token confirm.Token) string {
	return fmt.Sprintf("auth-svc:confirm:token-idx:%s", token)
}
