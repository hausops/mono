package testing

import (
	"context"
	"errors"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/rs/xid"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of user.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of user.Repository under test and teardown function.
//
// t must is needed to be passed to each newRepo so the cleanup runs for
// each subtest rather than once after the entire test suite finished.
func TestRepository(t *testing.T, newRepo func(t *testing.T) user.Repository) {
	ctx := context.Background()

	t.Run("Delete", func(t *testing.T) {
		repo := newRepo(t)
		uu := generateTestUsers(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, uu)

		// Delete a user that does not exist.
		_, err := repo.Delete(ctx, xid.New())
		if err != user.ErrNotFound {
			t.Error("Deleted user that does not exist")
		}

		// Delete a user.
		u := uu[1]
		deleted, err := repo.Delete(ctx, u.ID)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}
		if diff := cmp.Diff(u, deleted); diff != "" {
			t.Errorf("Delete returned incorrect user; (-want +got)\n%s", diff)
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
			u := uu[i]
			found, err := repo.FindByID(ctx, u.ID)
			if err != nil {
				t.Errorf("users[%d]: FindByID failed: %v", i, err)
			}
			if diff := cmp.Diff(u, found); diff != "" {
				t.Errorf("users[%d]: FindByID returned incorrect user; (-want +got)\n%s", i, diff)
			}
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		repo := newRepo(t)
		uu := generateTestUsers(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, uu)

		_, err := repo.FindByID(ctx, xid.New())
		if err != user.ErrNotFound {
			t.Errorf("FindByID(randomID) error = %v, want: ErrNotFound", err)
		}

		for i, u := range uu {
			found, err := repo.FindByID(ctx, u.ID)
			if err != nil {
				t.Errorf("users[%d]: FindByID failed: %v", i, err)
			}
			if diff := cmp.Diff(u, found); diff != "" {
				t.Errorf("users[%d]: FindByID returned incorrect user; (-want +got)\n%s", i, diff)
			}
		}
	})

	t.Run("FindByEmail", func(t *testing.T) {
		repo := newRepo(t)
		uu := generateTestUsers(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, uu)

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != user.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
		}

		for i, u := range uu {
			found, err := repo.FindByEmail(ctx, u.Email)
			if err != nil {
				t.Errorf("users[%d]: FindByEmail failed: %v", i, err)
			}
			if diff := cmp.Diff(u, found); diff != "" {
				t.Errorf("users[%d]: FindByEmail returned incorrect user; (-want +got)\n%s", i, diff)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new user", func(t *testing.T) {
			repo := newRepo(t)
			u := user.User{
				ID:    xid.New(),
				Email: mail.Address{Address: gofakeit.Email()},
			}

			inserted, err := repo.Upsert(ctx, u)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}
			if diff := cmp.Diff(u, inserted); diff != "" {
				t.Errorf("Upsert returned incorrect user; (-want +got)\n%s", diff)
			}

			// Find the user by ID after the insert.
			{
				found, err := repo.FindByID(ctx, u.ID)
				if err != nil {
					t.Errorf("FindByID failed: %v", err)
				}
				if diff := cmp.Diff(inserted, found); diff != "" {
					t.Errorf("FindByID returned incorrect user after insert; (-want +got)\n%s", diff)
				}
			}

			// Find the user by email after the insert.
			{
				found, err := repo.FindByEmail(ctx, u.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(inserted, found); diff != "" {
					t.Errorf("FindByEmail returned incorrect user after insert; (-want +got)\n%s", diff)
				}
			}
		})

		t.Run("Update with the same user info", func(t *testing.T) {
			repo := newRepo(t)
			u := user.User{
				ID:    xid.New(),
				Email: mail.Address{Address: gofakeit.Email()},
			}
			mustRepositoryUpsert(t, ctx, repo, u)

			updated, err := repo.Upsert(ctx, u)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}
			if diff := cmp.Diff(u, updated); diff != "" {
				t.Errorf("Upsert returned incorrect user; (-want +got)\n%s", diff)
			}
		})

		t.Run("Insert a new user with a duplicate email", func(t *testing.T) {
			repo := newRepo(t)
			u := user.User{
				ID:    xid.New(),
				Email: mail.Address{Address: gofakeit.Email()},
			}
			mustRepositoryUpsert(t, ctx, repo, u)

			_, err := repo.Upsert(ctx, user.User{
				ID:    xid.New(),
				Email: u.Email,
			})
			if !errors.Is(err, user.ErrEmailTaken) {
				t.Error("Upsert did not return ErrEmailTaken for duplicate email")
			}
		})

		t.Run("Update a user to use a duplicate email", func(t *testing.T) {
			repo := newRepo(t)
			uu := generateTestUsers(t, 3)
			mustRepositoryUpsertMany(t, ctx, repo, uu)

			u1 := uu[0]
			u2 := uu[1]

			// Update u2.Email to u1.Email
			_, err := repo.Upsert(ctx, user.User{
				ID:    u2.ID,
				Email: u1.Email,
			})
			if !errors.Is(err, user.ErrEmailTaken) {
				t.Error("Upsert did not return ErrEmailTaken for duplicate email")
			}
		})

		t.Run("Update an existing user with a different email", func(t *testing.T) {
			repo := newRepo(t)
			u := user.User{
				ID:    xid.New(),
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
			if diff := cmp.Diff(updated, up); diff != "" {
				t.Errorf("Upsert returned incorrect user; (-want +got)\n%s", diff)
			}

			// Find the user by ID after the update.
			{
				found, err := repo.FindByID(ctx, u.ID)
				if err != nil {
					t.Errorf("FindByID failed: %v", err)
				}
				if diff := cmp.Diff(updated, found); diff != "" {
					t.Errorf("FindByID returned incorrect user after insert; (-want +got)\n%s", diff)
				}
			}

			// Find the user by email after the update.
			{
				found, err := repo.FindByEmail(ctx, up.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(updated, found); diff != "" {
					t.Errorf("FindByEmail returned incorrect user after insert; (-want +got)\n%s", diff)
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
	if diff := cmp.Diff(u, inserted); diff != "" {
		t.Errorf("Upsert returned incorrect user; (-want +got)\n%s", diff)
	}
}

func mustRepositoryUpsertMany(
	t *testing.T,
	ctx context.Context,
	repo user.Repository,
	uu []user.User,
) {
	t.Helper()
	for i, u := range uu {
		inserted, err := repo.Upsert(ctx, u)
		if err != nil {
			t.Fatalf("users[%d]: Upsert failed: %v", i, err)
		}
		if diff := cmp.Diff(u, inserted); diff != "" {
			t.Errorf("users[%d]: Upsert returned incorrect user; (-want +got)\n%s", i, diff)
		}
	}
}
