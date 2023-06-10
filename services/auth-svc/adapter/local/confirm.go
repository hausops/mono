package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

// Pending confirmation

type pendingConfirmationRepository struct {
	byToken map[confirm.Token]confirm.Pending
}

func NewPendingConfirmationRepository() *pendingConfirmationRepository {
	return &pendingConfirmationRepository{
		byToken: make(map[confirm.Token]confirm.Pending),
	}
}

var _ confirm.PendingRepository = (*pendingConfirmationRepository)(nil)

func (r *pendingConfirmationRepository) DeleteByToken(ctx context.Context, token confirm.Token) (*confirm.Pending, error) {
	ver, ok := r.byToken[token]
	if !ok {
		return nil, confirm.ErrPendingNotFound
	}
	delete(r.byToken, token)
	return &ver, nil
}

func (r *pendingConfirmationRepository) FindByToken(ctx context.Context, token confirm.Token) (*confirm.Pending, error) {
	ver, ok := r.byToken[token]
	if !ok {
		return nil, confirm.ErrPendingNotFound
	}
	return &ver, nil
}

func (r *pendingConfirmationRepository) Upsert(ctx context.Context, pending confirm.Pending) error {
	r.byToken[pending.Token] = pending
	return nil
}

// Confirmed Email

type confirmedEmailRepository struct {
	byEmail map[mail.Address]struct{}
}

func NewConfirmedEmailRepository() *confirmedEmailRepository {
	return &confirmedEmailRepository{
		byEmail: make(map[mail.Address]struct{}),
	}
}

var _ confirm.ConfirmedEmailRepository = (*confirmedEmailRepository)(nil)

func (r *confirmedEmailRepository) Add(ctx context.Context, email mail.Address) error {
	_, ok := r.byEmail[email]
	if ok {
		return confirm.ErrEmailAlreadyExists
	}

	r.byEmail[email] = struct{}{}
	return nil
}

func (r *confirmedEmailRepository) Exist(ctx context.Context, email mail.Address) bool {
	_, ok := r.byEmail[email]
	return ok
}
