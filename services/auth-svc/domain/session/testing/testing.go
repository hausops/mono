// Package testing provides test clients and helpers for the session domain.
package testing

import (
	"net/mail"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/session"
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
	email := mail.Address{Address: gofakeit.Email()}
	return session.New(email, 15*time.Minute)
}
