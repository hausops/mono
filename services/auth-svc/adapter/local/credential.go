package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type credentialRepository struct {
	// Token is an index.
	byEmail map[mail.Address]user.ID
	// UserID is the primary key.
	byUserID map[user.ID]credential.Credential
}

func NewCredentialRepository() *credentialRepository {
	return &credentialRepository{
		byEmail:  make(map[mail.Address]user.ID),
		byUserID: make(map[user.ID]credential.Credential),
	}
}

var _ credential.Repository = (*credentialRepository)(nil)

func (r *credentialRepository) FindByEmail(_ context.Context, email mail.Address) (credential.Credential, error) {
	uid, ok := r.byEmail[email]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}

	cred, ok := r.byUserID[uid]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}
	return cred, nil
}

func (r *credentialRepository) FindByUserID(_ context.Context, uid user.ID) (credential.Credential, error) {
	cred, ok := r.byUserID[uid]
	if !ok {
		return credential.Credential{}, credential.ErrNotFound
	}
	return cred, nil
}

func (r *credentialRepository) Upsert(_ context.Context, cred credential.Credential) error {
	uid := cred.UserID

	// If updating, remove the previous token index for the record.
	if prev, ok := r.byUserID[uid]; ok {
		delete(r.byEmail, prev.Email)
	}

	r.byEmail[cred.Email] = uid
	r.byUserID[uid] = cred
	return nil
}
