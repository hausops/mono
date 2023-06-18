// Package testing provides test clients and helpers for the confirm domain.
package testing

import (
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

func generateTestRecords(t *testing.T, count int) []confirm.Record {
	t.Helper()
	records := make([]confirm.Record, count)
	for i := 0; i < len(records); i++ {
		token := confirm.GenerateToken()
		records[i] = confirm.Record{
			Email:       mail.Address{Address: gofakeit.Email()},
			Token:       token,
			IsConfirmed: false,
		}
	}
	return records
}
