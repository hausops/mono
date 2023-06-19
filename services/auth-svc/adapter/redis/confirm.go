package redis

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

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

func (r *confirmRepository) FindByEmail(ctx context.Context, email mail.Address) (confirm.Record, error) {
	primaryKey := r.primaryKey(email)

	var rec confirm.Record
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		// We need to check with Exists before calling HGetAll because,
		// when the key doesn't exist HGetAll returns nil error, instead of redis.Nil.
		// https://github.com/redis/go-redis/issues/1668
		n, err := tx.Exists(ctx, primaryKey).Result()
		if err != nil {
			return fmt.Errorf("redis.Exists(%s): %w", primaryKey, err)
		} else if n == 0 {
			return confirm.ErrNotFound
		}

		var saved confirmRecord
		err = tx.HGetAll(ctx, primaryKey).Scan(&saved)
		if err != nil {
			return fmt.Errorf("redis.HGetAll(%s): %w", primaryKey, err)
		}

		token, err := confirm.ParseToken(saved.Token)
		if err != nil {
			return err
		}

		rec = confirm.Record{
			Email:       email,
			IsConfirmed: saved.Confirmed,
			Token:       token,
		}
		return nil
	}, primaryKey)

	return rec, err
}

func (r *confirmRepository) FindByToken(ctx context.Context, token confirm.Token) (confirm.Record, error) {
	tokenKey := r.tokenKey(token)

	var rec confirm.Record
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		emailAddr, err := tx.Get(ctx, tokenKey).Result()
		switch {
		case errors.Is(err, redis.Nil):
			return confirm.ErrNotFound
		case err != nil:
			return fmt.Errorf("get email from token %s: %w", token, err)
		}

		email, err := mail.ParseAddress(emailAddr)
		if err != nil {
			return fmt.Errorf("parse email address: %w", err)
		}

		rec, err = r.FindByEmail(ctx, *email)
		if err != nil {
			return fmt.Errorf("FindByEmail(%s): %w", email.Address, err)
		}
		return nil
	}, tokenKey)

	return rec, err
}

func (r *confirmRepository) Upsert(ctx context.Context, rec confirm.Record) error {
	primaryKey := r.primaryKey(rec.Email)
	// Watch the primary key to detect changes by other clients
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		// Get the current token or empty string
		prevTokenStr, err := tx.HGet(ctx, primaryKey, "token").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, token): %w", primaryKey, err)
		}

		pipe := tx.Pipeline()
		pipe.HSet(ctx, primaryKey, confirmRecord{
			Confirmed: rec.IsConfirmed,
			Token:     rec.Token.String(),
		})

		// Remove the previous token index
		if prevTokenStr != "" {
			prevToken, err := confirm.ParseToken(prevTokenStr)
			if err != nil {
				return fmt.Errorf("parse previously stored token: %w", err)
			}
			pipe.Del(ctx, r.tokenKey(prevToken))
		}

		if !rec.Token.IsZero() {
			pipe.Set(ctx, r.tokenKey(rec.Token), rec.Email.Address, 0)
		}

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)
}

func (r *confirmRepository) primaryKey(email mail.Address) string {
	return fmt.Sprintf("auth-svc:confirm-repo:email:%s", email.Address)
}

func (r *confirmRepository) tokenKey(token confirm.Token) string {
	return fmt.Sprintf("auth-svc:confirm-repo:token-idx:%s", token)
}

// confirmRecord represents stored record data for a given key in redis.
type confirmRecord struct {
	Confirmed bool   `redis:"confirmed"`
	Token     string `redis:"token"`
}
