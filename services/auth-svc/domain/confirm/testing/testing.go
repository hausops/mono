// Package testing provides test clients and helpers for the confirm domain.
package testing

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

func generateTestRecord(t *testing.T, confirmed bool) confirm.Record {
	if confirmed {
		return confirm.Record{
			IsConfirmed: true,
			UserID:      user.NewID(),
		}
	}

	token := confirm.NewToken()
	return confirm.Record{
		IsConfirmed: false,
		Token:       token,
		UserID:      user.NewID(),
	}
}
