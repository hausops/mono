package redis

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/redis/go-redis/v9"
)

type credentialRepository struct {
	client *redis.Client
}

func NewCredentialRepository(client *redis.Client) *credentialRepository {
	return &credentialRepository{client: client}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(ctx context.Context, email mail.Address) (*credential.Credential, error) {
	k := r.withKeyPrefix(email.Address)
	password, err := r.client.Get(ctx, k).Bytes()
	switch {
	case errors.Is(err, redis.Nil):
		return nil, credential.ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("redis.Get(%s): %w", email.Address, err)
	}

	cred := credential.Credential{
		Email:    email,
		Password: password,
	}
	return &cred, nil
}

func (r *credentialRepository) Upsert(ctx context.Context, cred credential.Credential) error {
	k := r.withKeyPrefix(cred.Email.Address)
	v := cred.Password
	return r.client.Set(ctx, k, v, 0).Err()
}

func (r *credentialRepository) withKeyPrefix(key string) string {
	return fmt.Sprintf("auth-svc.repos.credential.%s", key)
}
