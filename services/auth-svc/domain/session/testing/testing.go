// Package testing provides test clients and helpers for the session domain.
package testing

import (
	"testing"
	"time"

	"github.com/hausops/mono/services/auth-svc/domain/session"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

func generateTestSessions(t *testing.T, count int) []session.Session {
	t.Helper()
	ss := make([]session.Session, count)
	for i := 0; i < len(ss); i++ {
		ss[i] = generateTestSession(t)
	}
	return ss
}

func generateTestSession(t *testing.T) session.Session {
	return session.New(user.NewID(), 15*time.Minute)
}
