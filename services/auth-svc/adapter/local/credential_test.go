package local_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	credtesting "github.com/hausops/mono/services/auth-svc/domain/credential/testing"
)

func TestUserRepository(t *testing.T) {
	credtesting.TestRepository(t, func(t *testing.T) credential.Repository {
		return local.NewCredentialRepository()
	})
}
