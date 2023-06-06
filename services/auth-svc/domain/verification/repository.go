package verification

import (
	"context"
)

type Repository interface {
	FindByToken(ctx context.Context, token Token) (*PendingVerification, error)

	Upsert(ctx context.Context, ver PendingVerification) error
}
