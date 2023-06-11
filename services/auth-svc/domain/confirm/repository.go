package confirm

import (
	"context"
	"net/mail"
)

// Repository interface declares the behavior this package needs
// to perists and retrieve data related to email confirmation.
type Repository interface {
	// FindByToken retrieves a record for the given email address.
	FindByEmail(context.Context, mail.Address) (Record, error)

	// FindByToken retrieves a record based on the given token.
	FindByToken(context.Context, Token) (Record, error)

	// Upsert inserts or updates an email confirmation record.
	Upsert(context.Context, Record) error
}
