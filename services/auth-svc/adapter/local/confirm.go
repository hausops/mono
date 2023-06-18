package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

type confirmRepository struct {
	// Email is the primary key.
	byEmail map[mail.Address]confirm.Record
	// Token is an index.
	byToken map[confirm.Token]mail.Address
}

func NewConfirmRepository() *confirmRepository {
	return &confirmRepository{
		byEmail: make(map[mail.Address]confirm.Record),
		byToken: make(map[confirm.Token]mail.Address),
	}
}

var _ confirm.Repository = (*confirmRepository)(nil)

func (r *confirmRepository) FindByEmail(_ context.Context, email mail.Address) (confirm.Record, error) {
	rec, ok := r.byEmail[email]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) FindByToken(_ context.Context, token confirm.Token) (confirm.Record, error) {
	email, ok := r.byToken[token]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}

	rec, ok := r.byEmail[email]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) Upsert(_ context.Context, rec confirm.Record) error {
	email := rec.Email

	// If updating, remove the previous token index for the record.
	if prev, ok := r.byEmail[email]; ok {
		delete(r.byToken, prev.Token)
	}

	r.byEmail[email] = rec
	if !rec.Token.IsZero() {
		r.byToken[rec.Token] = email
	}
	return nil
}
