package credential

import (
	"context"
	"net/mail"
)

type Repository interface {
	FindByEmail(ctx context.Context, email mail.Address) (*Credential, error)

	Upsert(ctx context.Context, cred Credential) error
}
