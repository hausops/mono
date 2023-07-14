package local

import (
	"context"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type confirmRepository struct {
	// Token is an index.
	byToken map[confirm.Token]user.ID
	// UserID is the primary key.
	byUserID map[user.ID]confirm.Record
}

func NewConfirmRepository() *confirmRepository {
	return &confirmRepository{
		byToken:  make(map[confirm.Token]user.ID),
		byUserID: make(map[user.ID]confirm.Record),
	}
}

var _ confirm.Repository = (*confirmRepository)(nil)

func (r *confirmRepository) FindByToken(_ context.Context, token confirm.Token) (confirm.Record, error) {
	uid, ok := r.byToken[token]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}

	rec, ok := r.byUserID[uid]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) FindByUserID(_ context.Context, uid user.ID) (confirm.Record, error) {
	rec, ok := r.byUserID[uid]
	if !ok {
		return confirm.Record{}, confirm.ErrNotFound
	}
	return rec, nil
}

func (r *confirmRepository) Upsert(_ context.Context, rec confirm.Record) error {
	uid := rec.UserID

	// If updating, remove the previous token index for the record.
	if prev, ok := r.byUserID[uid]; ok {
		delete(r.byToken, prev.Token)
	}

	r.byUserID[uid] = rec
	if !rec.Token.IsZero() {
		r.byToken[rec.Token] = uid
	}
	return nil
}
