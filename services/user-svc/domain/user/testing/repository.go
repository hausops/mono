package testing

import (
	"context"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of user.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of user.Repository under test.
func TestRepository(t *testing.T, newRepo func() user.Repository) {
	ctx := context.Background()

	t.Run("Delete", func(t *testing.T) {
		repo := newRepo()
		us := generateTestUsers(t, 3)
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
	})

	t.Run("FindByID", func(t *testing.T) {
		repo := newRepo()
		us := generateTestUsers(t, 3)
		for i, u := range us {
			if _, err := repo.Upsert(ctx, u); err != nil {
				t.Fatalf("Upsert users[%d] failed: %v", i, err)
			}
		}

		_, err := repo.FindByID(ctx, uuid.New())
		if err != user.ErrNotFound {
			t.Errorf("FindByID(randomID) error = %v, want: ErrNotFound", err)
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
	})

	t.Run("FindByEmail", func(t *testing.T) {
		repo := newRepo()
		us := generateTestUsers(t, 3)
		for i, u := range us {
			if _, err := repo.Upsert(ctx, u); err != nil {
				t.Fatalf("Upsert users[%d] failed: %v", i, err)
			}
		}

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != user.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
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
	})

	t.Run("Upsert", func(t *testing.T) {
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
			mustRepositoryUpsert(t, ctx, repo, u)

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
			mustRepositoryUpsert(t, ctx, repo, u)

			u2 := user.User{
				ID:    uuid.New(),
				Email: mail.Address{Address: gofakeit.Email()},
			}
			mustRepositoryUpsert(t, ctx, repo, u2)

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
			mustRepositoryUpsert(t, ctx, repo, u)

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
	})
}

func mustRepositoryUpsert(
	t *testing.T,
	ctx context.Context,
	repo user.Repository,
	u user.User,
) {
	t.Helper()
	inserted, err := repo.Upsert(ctx, u)
	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}
	if inserted != u {
		t.Fatal("Upsert returned incorrect user")
	}
}
