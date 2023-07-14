package confirm

import (
	"context"

	"github.com/hausops/mono/services/user-svc/domain/user"
)

// Repository interface declares the behavior this package needs
// to perists and retrieve data related to email confirmation.
type Repository interface {
	// FindByToken retrieves a record based on the given token.
	FindByToken(context.Context, Token) (Record, error)

	// FindByUserID retrieves a record for the given user ID.
	FindByUserID(context.Context, user.ID) (Record, error)

	// Upsert inserts or updates an email confirmation record.
	Upsert(context.Context, Record) error
}
