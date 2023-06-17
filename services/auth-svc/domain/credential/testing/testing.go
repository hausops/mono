// Package testing provides test clients and helpers for the credential domain.
package testing

import (
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
)

func generateTestCredentails(t *testing.T, count int) []credential.Credential {
	t.Helper()
	cc := make([]credential.Credential, count)
	for i := 0; i < len(cc); i++ {
		cc[i] = credential.Credential{
			Email:    mail.Address{Address: gofakeit.Email()},
			Password: generateTestPassword(t),
		}
	}
	return cc
}

func generateTestPassword(t *testing.T) []byte {
	t.Helper()
	len := gofakeit.Number(12, 20)
	return []byte(gofakeit.Password(true, true, true, true, false, len))
}
