package session

import (
	"time"

	"github.com/hausops/mono/services/user-svc/domain/user"
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

// IsExpires returns true when the elapsed time exceeds the sess.ExpireAt.
func (sess Session) IsExpired() bool {
	now := time.Now().UTC().Truncate(time.Second)
	return sess.ExpireAt.Before(now)
}
