package local_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/domain/session"
	sessiontesting "github.com/hausops/mono/services/auth-svc/domain/session/testing"
)

func TestSessionRepository(t *testing.T) {
	sessiontesting.TestRepository(t, func(t *testing.T) session.Repository {
		return local.NewSessionRepository()
	})
}
