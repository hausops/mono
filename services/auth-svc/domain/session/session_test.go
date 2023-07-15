package session_test

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/hausops/mono/services/auth-svc/domain/session"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

func TestSession_IsExpired(t *testing.T) {
	mockClock := clock.NewMock()

	expireAfter := 5 * time.Second
	sess := session.New(user.NewID(), expireAfter, session.WithClock(mockClock))

	if sess.IsExpired() {
		t.Error("sess.IsExpired() = true, immediately after creation")
	}

	mockClock.Add(expireAfter)
	if sess.IsExpired() {
		t.Error("sess.IsExpired() = true, before the elapsed time exceeds expireAfter")
	}

	mockClock.Add(time.Second)
	if !sess.IsExpired() {
		t.Error("sess.IsExpired() = false, after the elapsed time exceeds expireAfter")
	}
}
