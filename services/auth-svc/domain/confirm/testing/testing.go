// Package testing provides test clients and helpers for the confirm domain.
package testing

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/rs/xid"
)

func generateTestRecord(t *testing.T, confirmed bool) confirm.Record {
	if confirmed {
		return confirm.Record{
			IsConfirmed: true,
			UserID:      xid.New().String(),
		}
	}

	token := confirm.GenerateToken()
	return confirm.Record{
		IsConfirmed: false,
		Token:       token,
		UserID:      xid.New().String(),
	}
}
