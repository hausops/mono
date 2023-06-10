package verification

import (
	"context"
	"net/mail"
)

// PendingRepository interface declares the behavior this package needs
// to perists and retrieve data related to pending email verifications.
type PendingRepository interface {
	// DeleteByToken removes a pending verification record based on the given token.
	// It returns the deleted record if the deletion is successful.
	DeleteByToken(ctx context.Context, token Token) (*Pending, error)

	// FindByToken retrieves a pending verification record based on the given token.
	FindByToken(ctx context.Context, token Token) (*Pending, error)

	// Upsert inserts or updates a pending verification record.
	Upsert(ctx context.Context, pending Pending) error
}

// VerifiedRepository interface declares the behavior this package needs
// to perists and retrieve data related to verified email addresses.
type VerifiedRepository interface {
	// ExistByEmail checks whether the given email address is verified.
	// A nil error means the email is verified.
	ExistByEmail(ctx context.Context, email mail.Address) error

	// Upsert inserts or updates a verified email record.
	Upsert(ctx context.Context, email mail.Address) error
}
