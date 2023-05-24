package local_test

import (
	"context"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

// TODO: move this to a suite of contract tests exported by domain (user.Repository)
// so concrete implementations can run the suite to ensure the implementation
// conforms to the expected behavior.

func TestUserRepository_Delete(t *testing.T) {
	ctx := context.Background()
	repo := local.NewUserRepository()

	var users []user.User
	for i := 0; i < 3; i++ {
		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		users = append(users, u)
	}

	for i, u := range users {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	// Test deleting a user.
	u := users[1]
	deleted, err := repo.Delete(ctx, u.ID)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	if deleted != u {
		t.Error("Delete returned incorrect user")
	}

	// The deleted user is no longer found by ID.
	_, err = repo.FindByID(ctx, u.ID)
	if err != user.ErrNotFound {
		t.Error("Deleted user found by ID")
	}

	// The deleted user is no longer found by email.
	_, err = repo.FindByEmail(ctx, u.Email)
	if err != user.ErrNotFound {
		t.Error("Deleted user found by email")
	}

	// The other users still exist in the repository.
	for _, i := range []int{0, 2} {
		u := users[i]
		found, err := repo.FindByID(ctx, u.ID)
		if err != nil {
			t.Errorf("FindByID users[%d] failed: %v", i, err)
		}
		if found != u {
			t.Errorf("FindByID users[%d] returned incorrect user", i)
		}
	}
}
