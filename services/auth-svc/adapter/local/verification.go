package local

import (
	"context"

	"github.com/hausops/mono/services/auth-svc/domain/verification"
)

type verificationRepository struct {
	byToken map[verification.Token]verification.PendingVerification
}

func NewVerificationRepository() *verificationRepository {
	return &verificationRepository{
		byToken: make(map[verification.Token]verification.PendingVerification),
	}
}

var _ verification.Repository = (*verificationRepository)(nil)

func (r *verificationRepository) FindByToken(ctx context.Context, token verification.Token) (*verification.PendingVerification, error) {
	ver, ok := r.byToken[token]
	if !ok {
		return nil, verification.ErrNotFound
	}
	return &ver, nil
}

func (r *verificationRepository) Upsert(ctx context.Context, ver verification.PendingVerification) error {
	r.byToken[ver.Token] = ver
	return nil
}
