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

func NewCredentialRepository(c *redis.Client) *credentialRepository {
	return &credentialRepository{client: c}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(ctx context.Context, email mail.Address) (credential.Credential, error) {
	k := r.key(email)
	password, err := r.client.Get(ctx, k).Bytes()
	switch {
	case errors.Is(err, redis.Nil):
		return credential.Credential{}, credential.ErrNotFound
	case err != nil:
		return credential.Credential{}, fmt.Errorf("redis.Get(%s): %w", email.Address, err)
	}

	cred := credential.Credential{
		Email:    email,
		Password: password,
	}
	return cred, nil
}

func (r *credentialRepository) Upsert(ctx context.Context, cred credential.Credential) error {
	k := r.key(cred.Email)
	v := cred.Password
	return r.client.Set(ctx, k, v, 0).Err()
}

// key formats the primary key for storing a credential value in redis.
func (r *credentialRepository) key(email mail.Address) string {
	return fmt.Sprintf("auth-svc:credential:%s", email.Address)
}
