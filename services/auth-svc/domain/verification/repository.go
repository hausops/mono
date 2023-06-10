package verification

import (
	"context"
	"net/mail"
)

// PendingRepository interface declares the behavior this package needs
// to perists and retrieve data related to pending email verifications.
type PendingRepository interface {
	// DeleteByToken deletes the pending verification record associated
	// with the given token.
	// It returns the deleted record if the deletion is successful.
	// It returns ErrNotFound if no matching record is found for the token.
	DeleteByToken(ctx context.Context, token Token) (*Pending, error)

	// FindByToken retrieves a pending verification record based on the given token.
	FindByToken(ctx context.Context, token Token) (*Pending, error)

	// Upsert inserts or updates a pending verification record.
	Upsert(ctx context.Context, pending Pending) error
}

// VerifiedEmailRepository interface declares the behavior this package needs
// to perists and retrieve data related to verified email addresses.
type VerifiedEmailRepository interface {
	// Add inserts email to the repository.
	// If the email already exists, it returns ErrEmailAlreadyExists.
	Add(ctx context.Context, email mail.Address) error

	// Exist checks whether the given email address is in the repository.
	Exist(ctx context.Context, email mail.Address) bool
}
