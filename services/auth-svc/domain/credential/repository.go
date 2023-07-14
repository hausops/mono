package credential

import (
	"context"
	"net/mail"
)

// Repository interface declares the behavior this package needs to perists and
// retrieve data.
type Repository interface {
	// FindByEmail returns the credential with the given email, or an error
	// if the credential was not found.
	FindByEmail(context.Context, mail.Address) (Credential, error)

	// FindByUserID returns the credential with the given user ID, or an error
	// if the credential was not found.
	FindByUserID(ctx context.Context, userID string) (Credential, error)

	// Upsert adds cred to the repository if it does not exist, or replaces
	// the stored credential with the same email (without merging).
	Upsert(context.Context, Credential) error
}
