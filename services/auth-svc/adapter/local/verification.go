package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/verification"
)

// Pending Verification

type pendingVerificationRepository struct {
	byToken map[verification.Token]verification.Pending
}

func NewPendingVerificationRepository() *pendingVerificationRepository {
	return &pendingVerificationRepository{
		byToken: make(map[verification.Token]verification.Pending),
	}
}

var _ verification.PendingRepository = (*pendingVerificationRepository)(nil)

func (r *pendingVerificationRepository) DeleteByToken(ctx context.Context, token verification.Token) (*verification.Pending, error) {
	ver, ok := r.byToken[token]
	if !ok {
		return nil, verification.ErrPendingNotFound
	}
	delete(r.byToken, token)
	return &ver, nil
}

func (r *pendingVerificationRepository) FindByToken(ctx context.Context, token verification.Token) (*verification.Pending, error) {
	ver, ok := r.byToken[token]
	if !ok {
		return nil, verification.ErrPendingNotFound
	}
	return &ver, nil
}

func (r *pendingVerificationRepository) Upsert(ctx context.Context, pending verification.Pending) error {
	r.byToken[pending.Token] = pending
	return nil
}

// Verified Email

type verifiedEmailRepository struct {
	byEmail map[mail.Address]struct{}
}

func NewVerifiedEmailRepository() *verifiedEmailRepository {
	return &verifiedEmailRepository{
		byEmail: make(map[mail.Address]struct{}),
	}
}

var _ verification.VerifiedEmailRepository = (*verifiedEmailRepository)(nil)

func (r *verifiedEmailRepository) Add(ctx context.Context, email mail.Address) error {
	_, ok := r.byEmail[email]
	if ok {
		return verification.ErrEmailAlreadyExists
	}

	r.byEmail[email] = struct{}{}
	return nil
}

func (r *verifiedEmailRepository) Exist(ctx context.Context, email mail.Address) bool {
	_, ok := r.byEmail[email]
	return ok
}
