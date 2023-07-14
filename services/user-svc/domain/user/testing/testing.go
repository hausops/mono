// Package testing provide test clients and helpers for the user domain.
package testing

import (
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

func generateTestUsers(t *testing.T, count int) []user.User {
	t.Helper()
	uu := make([]user.User, count)
	for i := 0; i < len(uu); i++ {
		uu[i] = user.User{
			ID:    user.NewID(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
	}
	return uu
}
