package session

import (
	"net/mail"
	"time"

	"github.com/rs/xid"
)

type Session struct {
	AccessToken AccessToken
	Email       mail.Address
	ExpireAt    time.Time
}

func NewSession(email mail.Address, expireAfter time.Duration) Session {
	return Session{
		AccessToken: NewAccessToken(),
		Email:       email,
		ExpireAt:    time.Now().UTC().Add(expireAfter),
	}
}

func (sess Session) IsExpired() bool {
	return sess.ExpireAt.Before(time.Now().UTC())
}

type AccessToken xid.ID

func NewAccessToken() AccessToken {
	return AccessToken(xid.New())
}

func (at AccessToken) String() string {
	return xid.ID(at).String()
}

type TokenStore interface {
	// Store used to store a new token entry.
	// Store(ctx context.Context, token Token) error

	// Lookup used to get token entry by its signature.
	// Lookup(ctx context.Context, signature string) (Token, error)

	// Revoke used to delete token entry by its signature.
	// Revoke(ctx context.Context, signature string) error
}
