package redis

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/redis/go-redis/v9"
)

type credentialRepository struct {
	client *redis.Client
}

func NewCredentialRepository(c *redis.Client) *credentialRepository {
	return &credentialRepository{client: c}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(ctx context.Context, email mail.Address) (credential.Credential, error) {
	emailKey := r.emailKey(email)

	var cred credential.Credential
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		uidStr, err := tx.Get(ctx, emailKey).Result()
		switch {
		case errors.Is(err, redis.Nil):
			return credential.ErrNotFound
		case err != nil:
			return fmt.Errorf("get user ID from email %s: %w", email.Address, err)
		}

		uid, err := user.ParseID(uidStr)
		if err != nil {
			return fmt.Errorf("user.ParseID(%s): %w", uidStr, err)
		}

		cred, err = r.FindByUserID(ctx, uid)
		if err != nil {
			return fmt.Errorf("FindByUserID(%s): %w", uid, err)
		}
		return nil
	}, emailKey)

	return cred, err
}

func (r *credentialRepository) FindByUserID(ctx context.Context, uid user.ID) (credential.Credential, error) {
	primaryKey := r.primaryKey(uid)

	var cred credential.Credential
	err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		// We need to check with Exists before calling HGetAll because,
		// when the key doesn't exist HGetAll returns nil error, instead of redis.Nil.
		// https://github.com/redis/go-redis/issues/1668
		n, err := tx.Exists(ctx, primaryKey).Result()
		if err != nil {
			return fmt.Errorf("redis.Exists(%s): %w", primaryKey, err)
		} else if n == 0 {
			return credential.ErrNotFound
		}

		var saved credentialRedis
		err = tx.HGetAll(ctx, primaryKey).Scan(&saved)
		if err != nil {
			return fmt.Errorf("redis.HGetAll(%s): %w", primaryKey, err)
		}

		email, err := mail.ParseAddress(saved.Email)
		if err != nil {
			return fmt.Errorf("parse previously stored email %s: %w", saved.Email, err)
		}

		cred = credential.Credential{
			Email:    *email,
			Password: saved.Password,
			UserID:   uid,
		}
		return nil
	}, primaryKey)

	return cred, err
}

func (r *credentialRepository) Upsert(ctx context.Context, cred credential.Credential) error {
	primaryKey := r.primaryKey(cred.UserID)
	// Watch the primary key to detect changes by other clients
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		// Get the current email address or empty string
		prevEmailStr, err := tx.HGet(ctx, primaryKey, "email").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("redis.HGet(%s, email): %w", primaryKey, err)
		}

		pipe := tx.Pipeline()

		// Remove the previous email index
		if prevEmailStr != "" {
			prevEmail, err := mail.ParseAddress(prevEmailStr)
			if err != nil {
				return fmt.Errorf("parse previously stored email %s: %w", prevEmailStr, err)
			}
			pipe.Del(ctx, r.emailKey(*prevEmail))
		}

		pipe.HSet(ctx, primaryKey, credentialRedis{
			Email:    cred.Email.Address,
			Password: cred.Password,
		})

		pipe.Set(ctx, r.emailKey(cred.Email), cred.UserID.String(), 0)

		_, err = pipe.Exec(ctx)
		return err
	}, primaryKey)
}

// key formats the primary key for storing a credential value in redis.
func (r *credentialRepository) primaryKey(uid user.ID) string {
	return fmt.Sprintf("auth-svc:credential:%s", uid)
}

func (r *credentialRepository) emailKey(email mail.Address) string {
	return fmt.Sprintf("auth-svc:credential:email-idx:%s", email.Address)
}

// credentialRedis represents stored credential data for a given key in redis.
type credentialRedis struct {
	Email    string `redis:"email"`
	Password []byte `redis:"password"`
}
