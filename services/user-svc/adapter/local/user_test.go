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
	us := generateTestUsers(t)
	for i, u := range us {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	// Delete a user that does not exist.
	_, err := repo.Delete(ctx, uuid.New())
	if err != user.ErrNotFound {
		t.Error("Deleted user that does not exist")
	}

	// Delete a user.
	u := us[1]
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
		u := us[i]
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
	us := generateTestUsers(t)
	for i, u := range us {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	for i, u := range us {
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
	us := generateTestUsers(t)
	for i, u := range us {
		if _, err := repo.Upsert(ctx, u); err != nil {
			t.Fatalf("Upsert users[%d] failed: %v", i, err)
		}
	}

	for i, u := range us {
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

	t.Run("Insert a new user", func(t *testing.T) {
		repo := local.NewUserRepository()
		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}

		inserted, err := repo.Upsert(ctx, u)
		if err != nil {
			t.Errorf("Upsert failed: %v", err)
		}
		if inserted != u {
			t.Error("Upsert returned incorrect user")
		}

		// Find the user by ID after the insert.
		{
			found, err := repo.FindByID(ctx, u.ID)
			if err != nil {
				t.Errorf("FindByID failed: %v", err)
			}
			if found != inserted {
				t.Error("FindByID returned incorrect user after insert")
			}
		}

		// Find the user by email after the insert.
		{
			found, err := repo.FindByEmail(ctx, u.Email)
			if err != nil {
				t.Errorf("FindByEmail failed: %v", err)
			}
			if found != inserted {
				t.Error("FindByEmail returned incorrect user after insert")
			}
		}
	})

	t.Run("Update with the same user info", func(t *testing.T) {
		repo := local.NewUserRepository()
		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}

		updated, err := repo.Upsert(ctx, u)
		if err != nil {
			t.Errorf("Upsert failed: %v", err)
		}
		if updated != u {
			t.Error("Upsert returned incorrect user")
		}
	})

	t.Run("Insert a new user with a duplicate email", func(t *testing.T) {
		repo := local.NewUserRepository()
		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		mustUpsert(t, ctx, repo, u)

		_, err := repo.Upsert(ctx, user.User{
			ID:    uuid.New(),
			Email: u.Email,
		})
		if err != user.ErrEmailAlreadyUsed {
			t.Error("Upsert did not return ErrEmailAlreadyUsed for duplicate email")
		}
	})

	t.Run("Update a user to use a duplicate email", func(t *testing.T) {
		repo := local.NewUserRepository()

		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		mustUpsert(t, ctx, repo, u)

		u2 := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		mustUpsert(t, ctx, repo, u2)

		// Update u2.Email to u.Email
		_, err := repo.Upsert(ctx, user.User{
			ID:    u2.ID,
			Email: u.Email,
		})
		if err != user.ErrEmailAlreadyUsed {
			t.Error("Upsert did not return ErrEmailAlreadyUsed for duplicate email")
		}
	})

	t.Run("Update an existing user with a different email", func(t *testing.T) {
		repo := local.NewUserRepository()
		u := user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
		mustUpsert(t, ctx, repo, u)

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
			if found != updated {
				t.Error("FindByID returned incorrect user after update")
			}
		}

		// Find the user by email after the update.
		{
			found, err := repo.FindByEmail(ctx, up.Email)
			if err != nil {
				t.Errorf("FindByEmail failed: %v", err)
			}
			if found != updated {
				t.Error("FindByEmail returned incorrect user after update")
			}
		}

		// Find the user by the _old_ email after the update.
		{
			_, err := repo.FindByEmail(ctx, u.Email)
			if err != user.ErrNotFound {
				t.Error("Updated user found by the old email")
			}
		}
	})
}

func mustUpsert(t *testing.T, ctx context.Context, repo user.Repository, u user.User) {
	t.Helper()
	inserted, err := repo.Upsert(ctx, u)
	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}
	if inserted != u {
		t.Fatal("Upsert returned incorrect user")
	}
}

func generateTestUsers(t *testing.T) []user.User {
	t.Helper()
	us := make([]user.User, 3)
	for i := 0; i < len(us); i++ {
		us[i] = user.User{
			ID:    uuid.New(),
			Email: mail.Address{Address: gofakeit.Email()},
		}
	}
	return us
}
