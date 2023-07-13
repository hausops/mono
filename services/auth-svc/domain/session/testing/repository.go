package testing

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/auth-svc/domain/session"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of session.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of session.Repository under test.
func TestRepository(t *testing.T, newRepo func(t *testing.T) session.Repository) {
	ctx := context.Background()

	t.Run("DeleteByAccessToken", func(t *testing.T) {
		repo := newRepo(t)
		sessions := generateTestSessions(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, sessions)

		// Delete a session that does not exist.
		err := repo.DeleteByAccessToken(ctx, generateTestSession(t).AccessToken)
		if err != session.ErrNotFound {
			t.Errorf("Deleted a session that does not exist; error = %v, want: ErrNotFound", err)
		}

		// Delete a session.
		sess := sessions[1]
		err = repo.DeleteByAccessToken(ctx, sess.AccessToken)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}

		// The deleted session is no longer found by access token.
		_, err = repo.FindByAccessToken(ctx, sess.AccessToken)
		if err != session.ErrNotFound {
			t.Error("Deleted session found by access token")
		}

		// The deleted session is no longer found by user ID.
		_, err = repo.FindByUserID(ctx, sess.UserID)
		if err != session.ErrNotFound {
			t.Error("Deleted session found by user ID")
		}

		// The other sessions still exist in the repository.
		for _, i := range []int{0, 2} {
			sess := sessions[i]

			found, err := repo.FindByAccessToken(ctx, sess.AccessToken)
			if err != nil {
				t.Errorf("sessions[%d]: FindByAccessToken failed: %v", i, err)
			}
			if diff := cmp.Diff(sess, found); diff != "" {
				t.Errorf("sessions[%d]: FindByAccessToken returned incorrect session; (-want +got)\n%s", i, diff)
			}

			found, err = repo.FindByUserID(ctx, sess.UserID)
			if err != nil {
				t.Errorf("sessions[%d]: FindByUserID failed: %v", i, err)
			}
			if diff := cmp.Diff(sess, found); diff != "" {
				t.Errorf("sessions[%d]: FindByUserID returned incorrect session; (-want +got)\n%s", i, diff)
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
				t.Errorf("sessions[%d]: FindByAccessToken failed: %v", i, err)
			}
			if diff := cmp.Diff(sess, found); diff != "" {
				t.Errorf("sessions[%d]: FindByAccessToken returned incorrect session; (-want +got)\n%s", i, diff)
			}
		}
	})

	t.Run("FindByUserID", func(t *testing.T) {
		repo := newRepo(t)
		sessions := generateTestSessions(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, sessions)

		_, err := repo.FindByUserID(ctx, generateTestSession(t).UserID)
		if err != session.ErrNotFound {
			t.Errorf("FindByUserID(randomUserID) error = %v, want: ErrNotFound", err)
		}

		for i, sess := range sessions {
			found, err := repo.FindByUserID(ctx, sess.UserID)
			if err != nil {
				t.Errorf("sessions[%d]: FindByUserID failed: %v", i, err)
			}
			if diff := cmp.Diff(sess, found); diff != "" {
				t.Errorf("sessions[%d]: FindByUserID returned incorrect session; (-want +got)\n%s", i, diff)
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

			// Find the session by access token after the insert.
			{
				found, err := repo.FindByAccessToken(ctx, sess.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if diff := cmp.Diff(sess, found); diff != "" {
					t.Errorf("FindByAccessToken returned incorrect session; (-want +got)\n%s", diff)
				}
			}

			// Find the session by user ID after the insert.
			{
				found, err := repo.FindByUserID(ctx, sess.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(sess, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect session; (-want +got)\n%s", diff)
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

			// Find the session by access token after the insert.
			{
				found, err := repo.FindByAccessToken(ctx, sess.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if diff := cmp.Diff(sess, found); diff != "" {
					t.Errorf("FindByAccessToken returned incorrect session; (-want +got)\n%s", diff)
				}
			}

			// Find the session by user ID after the insert.
			{
				found, err := repo.FindByUserID(ctx, sess.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(sess, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect session; (-want +got)\n%s", diff)
				}
			}
		})

		// Same user ID, different access token i.e. generating a new session
		t.Run("Update an existing session with a different access token", func(t *testing.T) {
			repo := newRepo(t)
			sess := generateTestSession(t)
			mustRepositoryUpsert(t, ctx, repo, sess)

			up := session.New(sess.UserID, 30*time.Minute)
			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the session by token after the update.
			{
				found, err := repo.FindByAccessToken(ctx, up.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByAccessToken returned incorrect session; (-want +got)\n%s", diff)
				}
			}

			// Find the session by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect session; (-want +got)\n%s", diff)
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

		// This handles an edge case that is unlikely to occur in normal operation.
		// However, we still handle it to ensure proper index updating and maintain
		// the desired data invariant.
		t.Run("Updating with the same access token but different user ID", func(t *testing.T) {
			repo := newRepo(t)
			sess := generateTestSession(t)
			mustRepositoryUpsert(t, ctx, repo, sess)

			up := session.Session{
				AccessToken: sess.AccessToken,
				ExpireAt:    sess.ExpireAt,
				UserID:      generateTestSession(t).UserID,
			}
			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the session by token after the update.
			{
				found, err := repo.FindByAccessToken(ctx, up.AccessToken)
				if err != nil {
					t.Errorf("FindByAccessToken failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByAccessToken returned incorrect session; (-want +got)\n%s", diff)
				}
			}

			// Find the session by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect session; (-want +got)\n%s", diff)
				}
			}

			// Find the session by the _old_ user ID after the update.
			{
				_, err := repo.FindByUserID(ctx, sess.UserID)
				if err != session.ErrNotFound {
					t.Error("Updated session found by the old user ID")
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
			t.Fatalf("sessions[%d]: Upsert failed: %v", i, err)
		}
	}
}
