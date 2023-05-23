package user

import (
	"context"
	"net/mail"

	"github.com/google/uuid"
)

// Repository interface declares the behavior this package needs to perists and
// retrieve data.
type Repository interface {
	// Delete removes the user with the given id and returns
	// the deleted user, or an error if the user was not found.
	Delete(ctx context.Context, id uuid.UUID) (User, error)

	// FindByID returns the user with the given id, or an error
	// if the user was not found.
	FindByID(ctx context.Context, id uuid.UUID) (User, error)

	// FindByEmail returns the user with the given email, or an error
	// if the user was not found.
	FindByEmail(ctx context.Context, email mail.Address) (User, error)

	// Upsert adds u to the repository if it does not exist, or replaces
	// the stored user with the same ID (without merging).
	Upsert(ctx context.Context, u User) (User, error)
}
