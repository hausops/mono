// Package testing provides test clients and helpers for the confirm domain.
package testing

import (
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

func generateTestRecord(t *testing.T, confirmed bool) confirm.Record {
	email := mail.Address{Address: gofakeit.Email()}

	if confirmed {
		return confirm.Record{
			Email:       email,
			IsConfirmed: true,
		}
	}

	token := confirm.GenerateToken()
	return confirm.Record{
		Email:       email,
		Token:       token,
		IsConfirmed: false,
	}
}
