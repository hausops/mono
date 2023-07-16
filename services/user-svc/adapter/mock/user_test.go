package mock_test

import (
	"testing"

	"github.com/hausops/mono/services/user-svc/adapter/mock"
	"github.com/hausops/mono/services/user-svc/domain/user"
	usertesting "github.com/hausops/mono/services/user-svc/domain/user/testing"
)

func TestUserRepository(t *testing.T) {
	usertesting.TestRepository(t, func(_ *testing.T) user.Repository {
		return mock.NewUserRepository()
	})
}
