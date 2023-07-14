package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
)

type credentialRepository struct {
	// Token is an index.
	byEmail map[mail.Address]string
	// UserID is the primary key.
	byUserID map[string]credential.Credential
}

func NewCredentialRepository() *credentialRepository {
	return &credentialRepository{
		byEmail:  make(map[mail.Address]string),
		byUserID: make(map[string]credential.Credential),
	}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(_ context.Context, email mail.Address) (credential.Credential, error) {
	userID, ok := r.byEmail[email]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}

	cred, ok := r.byUserID[userID]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}
	return cred, nil
}

func (r *credentialRepository) FindByUserID(_ context.Context, userID string) (credential.Credential, error) {
	cred, ok := r.byUserID[userID]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}
	return cred, nil
}

func (r *credentialRepository) Upsert(_ context.Context, cred credential.Credential) error {
	userID := cred.UserID

	// If updating, remove the previous token index for the record.
	if prev, ok := r.byUserID[userID]; ok {
		delete(r.byEmail, prev.Email)
	}

	r.byEmail[cred.Email] = userID
	r.byUserID[userID] = cred
	return nil
}
