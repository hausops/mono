package local_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	confirmtesting "github.com/hausops/mono/services/auth-svc/domain/confirm/testing"
)

func TestConfirmRepository(t *testing.T) {
	confirmtesting.TestRepository(t, func(t *testing.T) confirm.Repository {
		return local.NewConfirmRepository()
	})
}
