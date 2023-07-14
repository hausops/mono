package local

import (
	"context"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

type confirmRepository struct {
	// Token is an index.
	byToken map[confirm.Token]string
	// UserID is the primary key.
	byUserID map[string]confirm.Record
}

func NewConfirmRepository() *confirmRepository {
	return &confirmRepository{
		byToken:  make(map[confirm.Token]string),
		byUserID: make(map[string]confirm.Record),
	}
}

var _ confirm.Repository = (*confirmRepository)(nil)

func (r *confirmRepository) FindByToken(_ context.Context, token confirm.Token) (confirm.Record, error) {
	userID, ok := r.byToken[token]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}

	rec, ok := r.byUserID[userID]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) FindByUserID(_ context.Context, userID string) (confirm.Record, error) {
	rec, ok := r.byUserID[userID]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) Upsert(_ context.Context, rec confirm.Record) error {
	userID := rec.UserID

	// If updating, remove the previous token index for the record.
	if prev, ok := r.byUserID[userID]; ok {
		delete(r.byToken, prev.Token)
	}

	r.byUserID[userID] = rec
	if !rec.Token.IsZero() {
		r.byToken[rec.Token] = userID
	}
	return nil
}
