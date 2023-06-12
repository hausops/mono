package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
)

type credentialRepository struct {
	byEmail map[mail.Address]credential.Credential
}

func NewCredentialRepository() *credentialRepository {
	return &credentialRepository{
		byEmail: make(map[mail.Address]credential.Credential),
	}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(ctx context.Context, email mail.Address) (*credential.Credential, error) {
	cred, ok := r.byEmail[email]
	if !ok {
		return nil, credential.ErrNotFound
	}
	return &cred, nil
}

func (r *credentialRepository) Upsert(ctx context.Context, cred credential.Credential) error {
	r.byEmail[cred.Email] = cred
	return nil
}
