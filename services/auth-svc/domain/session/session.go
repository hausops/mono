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

func ParseAccessToken(b []byte) (AccessToken, error) {
	id, err := xid.FromBytes(b)
	if err != nil {
		return AccessToken{}, ErrInvalidToken
	}
	return AccessToken(id), nil
}
