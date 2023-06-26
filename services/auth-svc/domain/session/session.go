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

// New constructs a new Session with a new random access token
// with the expiration based on expireAfter.
func New(email mail.Address, expireAfter time.Duration) Session {
	now := time.Now().UTC().Truncate(time.Second)
	return Session{
		AccessToken: NewAccessToken(),
		Email:       email,
		ExpireAt:    now.Add(expireAfter.Truncate(time.Second)),
	}
}

func (sess Session) IsExpired() bool {
	now := time.Now().UTC().Truncate(time.Second)
	return sess.ExpireAt.Before(now)
}

type AccessToken xid.ID

func NewAccessToken() AccessToken {
	return AccessToken(xid.New())
}

func (at AccessToken) String() string {
	return xid.ID(at).String()
}

func ParseAccessToken(s string) (AccessToken, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return AccessToken{}, ErrInvalidToken
	}
	return AccessToken(id), nil
}
