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
	users := generateTestUsers(t)
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

func TestUserRepository_FindByID(t *testing.T) {
	ctx := context.Background()
	repo := local.NewUserRepository()
	users := generateTestUsers(t)
	for i, u := range users {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	for i, u := range users {
		found, err := repo.FindByID(ctx, u.ID)
		if err != nil {
			t.Errorf("FindByID users[%d] failed: %v", i, err)
		}
		if found != u {
			t.Errorf("FindByID users[%d] returned incorrect user", i)
		}
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	ctx := context.Background()
	repo := local.NewUserRepository()
	users := generateTestUsers(t)
	for i, u := range users {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	for i, u := range users {
		found, err := repo.FindByEmail(ctx, u.Email)
		if err != nil {
			t.Errorf("FindByEmail users[%d] failed: %v", i, err)
		}
		if found != u {
			t.Errorf("FindByEmail users[%d] returned incorrect user", i)
		}
	}
}

func TestUserRepository_Upsert(t *testing.T) {
	ctx := context.Background()
	repo := local.NewUserRepository()

	u := user.User{
		ID:    uuid.New(),
		Email: mail.Address{Address: gofakeit.Email()},
	}

	// Insert a new user.
	{
		inserted, err := repo.Upsert(ctx, u)
		if err != nil {
			t.Errorf("Upsert failed: %v", err)
		}
		if inserted != u {
			t.Error("Upsert returned incorrect user")
		}
	}

	// Update with the same user info.
	{
		updated, err := repo.Upsert(ctx, u)
		if err != nil {
			t.Errorf("Upsert failed: %v", err)
		}
		if updated != u {
			t.Error("Upsert returned incorrect user")
		}
	}

	// Insert a new user with a duplicate email.
	{
		_, err := repo.Upsert(ctx, user.User{
			ID:    uuid.New(),
			Email: u.Email,
		})
		if err != user.ErrEmailAlreadyUsed {
			t.Error("Upsert did not return ErrEmailAlreadyUsed for duplicate email")
		}
	}

	// Update a user to use a duplicate email.
	{
		// Setup: add a new user with a different email first.
		u2 := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		if _, err := repo.Upsert(ctx, u2); err != nil {
			t.Fatalf("Upsert user2 failed: %v", err)
		}

		_, err := repo.Upsert(ctx, user.User{
			ID:    u2.ID,
			Email: u.Email,
		})
		if err != user.ErrEmailAlreadyUsed {
			t.Error("Upsert did not return ErrEmailAlreadyUsed for duplicate email")
		}
	}

	// Update an existing user with a different email.
	up := user.User{
		ID:    u.ID,
		Email: mail.Address{Address: gofakeit.Email()},
	}

	updated, err := repo.Upsert(ctx, up)
	if err != nil {
		t.Errorf("Upsert failed: %v", err)
	}
	if updated != up {
		t.Error("Upsert returned incorrect user")
	}

	// Find the user by ID after the update.
	{
		found, err := repo.FindByID(ctx, u.ID)
		if err != nil {
			t.Errorf("FindByID failed: %v", err)
		}
		if found != up {
			t.Error("FindByID returned incorrect user after update")
		}
	}

	// Find the user by email after the update.
	{
		found, err := repo.FindByEmail(ctx, up.Email)
		if err != nil {
			t.Errorf("FindByEmail failed: %v", err)
		}
		if found != up {
			t.Error("FindByEmail returned incorrect user after update")
		}
	}
}

func generateTestUsers(t *testing.T) []user.User {
	t.Helper()
	users := make([]user.User, 3)
	for i := 0; i < len(users); i++ {
		users[i] = user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
	}
	return users
}
