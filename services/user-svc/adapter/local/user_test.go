package local_test

import (
	"testing"

	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/domain/user"
	usertesting "github.com/hausops/mono/services/user-svc/domain/user/testing"
)

func TestUserRepository(t *testing.T) {
	usertesting.TestRepository(t, func() user.Repository {
		return local.NewUserRepository()
	})
}
