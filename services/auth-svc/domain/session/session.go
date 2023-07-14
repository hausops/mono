package session

import (
	"time"

	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/rs/xid"
)

type Session struct {
	AccessToken AccessToken
	ExpireAt    time.Time
	UserID      user.ID
}

// New constructs a new Session with a new random access token
// with the expiration based on expireAfter.
func New(uid user.ID, expireAfter time.Duration) Session {
	now := time.Now().UTC().Truncate(time.Second)
	return Session{
		AccessToken: NewAccessToken(),
		ExpireAt:    now.Add(expireAfter.Truncate(time.Second)),
		UserID:      uid,
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
