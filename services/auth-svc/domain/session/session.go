package session

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type Session struct {
	AccessToken AccessToken
	ExpireAt    time.Time
	UserID      user.ID

	// clock is configurable for testing
	clock clock.Clock
}

// New constructs a new Session with a new random access token
// with the expiration based on expireAfter.
func New(uid user.ID, expireAfter time.Duration, options ...SessionOption) Session {
	sess := Session{
		AccessToken: NewAccessToken(),
		UserID:      uid,
	}

	// Apply options
	for _, opt := range options {
		opt(&sess)
	}

	// Use real clock if none is set.
	if sess.clock == nil {
		sess.clock = clock.New()
	}

	// Set ExpireAt after we know which clock to use.
	now := sess.clock.Now().UTC().Truncate(time.Second)
	sess.ExpireAt = now.Add(expireAfter.Truncate(time.Second))

	return sess
}

// SessionOption configures sess.
type SessionOption func(sess *Session)

// WithClock returns a SessionOption that sets clock on sess.
func WithClock(clock clock.Clock) SessionOption {
	return func(sess *Session) { sess.clock = clock }
}

// IsExpires returns true when the elapsed time exceeds the sess.ExpireAt.
func (sess Session) IsExpired() bool {
	now := sess.clock.Now().UTC().Truncate(time.Second)
	return sess.ExpireAt.Before(now)
}
