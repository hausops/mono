// Package testing provides test clients and helpers for the credential domain.
package testing

import (
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/rs/xid"
)

func generateTestCredentials(t *testing.T, count int) []credential.Credential {
	t.Helper()
	cc := make([]credential.Credential, count)
	for i := 0; i < len(cc); i++ {
		cc[i] = generateTestCredential(t)
	}
	return cc
}

func generateTestCredential(t *testing.T) credential.Credential {
	t.Helper()
	return credential.Credential{
		Email:    mail.Address{Address: gofakeit.Email()},
		Password: generateTestPassword(t),
		UserID:   xid.New().String(),
	}
}

func generateTestPassword(t *testing.T) []byte {
	t.Helper()
	len := gofakeit.Number(12, 20)
	return []byte(gofakeit.Password(true, true, true, true, false, len))
}
