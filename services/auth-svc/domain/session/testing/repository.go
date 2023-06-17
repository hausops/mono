package testing

import (
	"context"
	"net/mail"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/domain/session"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of session.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of session.Repository under test.
func TestRepository(t *testing.T, newRepo func(t *testing.T) session.Repository) {
	ctx := context.Background()

	t.Run("DeleteByEmail", func(t *testing.T) {
		repo := newRepo(t)
		sessions := generateTestSessions(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, sessions)

		// Delete a session that does not exist.
		_, err := repo.DeleteByEmail(ctx, generateTestSession(t).Email)
		if err != session.ErrNotFound {
			t.Error("Deleted a session that does not exist")
		}

		// Delete a session.
		sess := sessions[1]
		deleted, err := repo.DeleteByEmail(ctx, sess.Email)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}
		if deleted != sess {
			t.Error("Delete returned incorrect session")
		}

		// The deleted session is no longer found by email.
		_, err = repo.FindByEmail(ctx, deleted.Email)
		if err != session.ErrNotFound {
			t.Error("Deleted session found by email")
		}

		// The deleted session is no longer found by access token.
		_, err = repo.FindByAccessToken(ctx, deleted.AccessToken)
		if err != session.ErrNotFound {
			t.Error("Deleted session found by ID")
		}

		// The other sessions still exist in the repository.
		for _, i := range []int{0, 2} {
			sess := sessions[i]

			found, err := repo.FindByEmail(ctx, sess.Email)
			if err != nil {
				t.Errorf("FindByEmail sessions[%d] failed: %v", i, err)
			}
			if found != sess {
				t.Errorf("FindByEmail sessions[%d] returned incorrect session", i)
			}

			found, err = repo.FindByAccessToken(ctx, sess.AccessToken)
			if err != nil {
				t.Errorf("FindByAccessToken sessions[%d] failed: %v", i, err)
			}
			if found != sess {
				t.Errorf("FindByAccessToken sessions[%d] returned incorrect session", i)
			}
		}
	})

	t.Run("FindByEmail", func(t *testing.T) {
		repo := newRepo(t)
		sessions := generateTestSessions(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, sessions)

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != session.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range sessions {
			found, err := repo.FindByEmail(ctx, rec.Email)
			if err != nil {
				t.Errorf("FindByEmail sessions[%d] failed: %v", i, err)
			}
			if found != rec {
				t.Errorf("FindByEmail sessions[%d] returned incorrect session", i)
			}
		}
	})

	t.Run("FindByAccessToken", func(t *testing.T) {
		repo := newRepo(t)
		sessions := generateTestSessions(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, sessions)

		_, err := repo.FindByAccessToken(ctx, session.NewAccessToken())
		if err != session.ErrNotFound {
			t.Errorf("FindByAccessToken(randomToken) error = %v, want: ErrNotFound", err)
		}

		for i, sess := range sessions {
			found, err := repo.FindByAccessToken(ctx, sess.AccessToken)
			if err != nil {
				t.Errorf("FindByAccessToken sessions[%d] failed: %v", i, err)
			}
			if found != sess {
				t.Errorf("FindByAccessToken sessions[%d] returned incorrect session", i)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new session", func(t *testing.T) {
			repo := newRepo(t)
			sess := generateTestSession(t)

			err := repo.Upsert(ctx, sess)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the session by email after the insert.
			{
				found, err := repo.FindByEmail(ctx, sess.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if found != sess {
					t.Error("FindByEmail returned incorrect session after insert")
				}
			}

			// Find the session by access token after the insert.
			{
				found, err := repo.FindByAccessToken(ctx, sess.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if found != sess {
					t.Error("FindByAccessToken returned incorrect session after insert")
				}
			}
		})

		t.Run("Update with the same session", func(t *testing.T) {
			repo := newRepo(t)
			sess := generateTestSession(t)
			mustRepositoryUpsert(t, ctx, repo, sess)

			err := repo.Upsert(ctx, sess)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			found, err := repo.FindByEmail(ctx, sess.Email)
			if err != nil {
				t.Errorf("FindByEmail failed: %v", err)
			}
			if found != sess {
				t.Error("FindByEmail returned incorrect session")
			}
		})

		t.Run("Update an existing session with a different access token", func(t *testing.T) {
			repo := newRepo(t)
			sess := generateTestSession(t)
			mustRepositoryUpsert(t, ctx, repo, sess)

			up := session.New(sess.Email, 30*time.Minute)
			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the session by email after the update.
			{
				found, err := repo.FindByEmail(ctx, sess.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if found != up {
					t.Error("FindByEmail returned incorrect session after update")
				}
			}

			// Find the session by token after the update.
			{
				found, err := repo.FindByAccessToken(ctx, up.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if found != up {
					t.Error("FindByAccessToken returned incorrect session after update")
				}
			}

			// Find the session by the _old_ token after the update.
			{
				_, err := repo.FindByAccessToken(ctx, sess.AccessToken)
				if err != session.ErrNotFound {
					t.Error("Updated session found by the old token")
				}
			}
		})
	})
}

func mustRepositoryUpsert(
	t *testing.T,
	ctx context.Context,
	repo session.Repository,
	sess session.Session,
) {
	t.Helper()
	if err := repo.Upsert(ctx, sess); err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}
}

func mustRepositoryUpsertMany(
	t *testing.T,
	ctx context.Context,
	repo session.Repository,
	sessions []session.Session,
) {
	t.Helper()
	for i, sess := range sessions {
		if err := repo.Upsert(ctx, sess); err != nil {
			t.Fatalf("Upsert sessions[%d] failed: %v", i, err)
		}
	}
}
