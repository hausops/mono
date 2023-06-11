package session

import (
	"context"
	"net/mail"
)

type Repository interface {
	DeleteByEmail(context.Context, mail.Address) (Session, error)

	FindByAccessToken(context.Context, AccessToken) (Session, error)

	FindByEmail(context.Context, mail.Address) (Session, error)

	Upsert(context.Context, Session) error
}
