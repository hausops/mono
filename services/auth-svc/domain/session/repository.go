package session

import (
	"context"
	"net/mail"
)

type Repository interface {
	DeleteByAccessToken(context.Context, AccessToken) error

	FindByAccessToken(context.Context, AccessToken) (Session, error)

	FindByEmail(context.Context, mail.Address) (Session, error)

	Upsert(context.Context, Session) error
}
