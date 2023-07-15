package testing

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of confirm.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of confirm.Repository under test.
func TestRepository(t *testing.T, newRepo func(t *testing.T) confirm.Repository) {
	ctx := context.Background()

	t.Run("FindByToken", func(t *testing.T) {
		repo := newRepo(t)
		records := []confirm.Record{
			generateTestRecord(t, false),
			generateTestRecord(t, true),
			generateTestRecord(t, false),
			generateTestRecord(t, true),
		}
		mustRepositoryUpsertMany(t, ctx, repo, records)

		_, err := repo.FindByToken(ctx, confirm.NewToken())
		if err != confirm.ErrNotFound {
			t.Errorf("FindByToken(randomToken) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			found, err := repo.FindByToken(ctx, rec.Token)

			// Find confirmed records
			if rec.IsConfirmed {
				prefix := fmt.Sprintf("records[%d] (confirmed)", i)
				if err != confirm.ErrNotFound {
					t.Errorf("%s: FindByToken() error = %v, want: ErrNotFound", prefix, err)
				}
				continue
			}

			// Find unconfirmed records
			prefix := fmt.Sprintf("records[%d] (unconfirmed)", i)
			if err != nil {
				t.Errorf("%s: FindByToken failed: %v", prefix, err)
			}
			if diff := cmp.Diff(rec, found); diff != "" {
				t.Errorf("%s: FindByToken returned incorrect record; (-want +got)\n%s", prefix, diff)
			}
		}
	})

	t.Run("FindByUserID", func(t *testing.T) {
		repo := newRepo(t)
		records := []confirm.Record{
			generateTestRecord(t, false),
			generateTestRecord(t, true),
			generateTestRecord(t, false),
		}
		mustRepositoryUpsertMany(t, ctx, repo, records)

		_, err := repo.FindByUserID(ctx, user.NewID())
		if err != confirm.ErrNotFound {
			t.Errorf("FindByUserID(randomUserID) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			found, err := repo.FindByUserID(ctx, rec.UserID)
			if err != nil {
				t.Errorf("records[%d]: FindByUserID failed: %v", i, err)
			}
			if diff := cmp.Diff(rec, found); diff != "" {
				t.Errorf("records[%d]: FindByUserID returned incorrect record; (-want +got)\n%s", i, diff)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new record", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.NewToken()
			rec := confirm.Record{
				IsConfirmed: false,
				Token:       token,
				UserID:      user.NewID(),
			}

			err := repo.Upsert(ctx, rec)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by user ID after the insert.
			{
				found, err := repo.FindByUserID(ctx, rec.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(rec, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect record; (-want +got)\n%s", diff)
				}
			}

			// Find the record by token after the insert.
			{
				found, err := repo.FindByToken(ctx, token)
				if err != nil {
					t.Errorf("FindByToken failed: %v", err)
				}
				if diff := cmp.Diff(rec, found); diff != "" {
					t.Errorf("FindByToken returned incorrect record; (-want +got)\n%s", diff)
				}
			}
		})

		t.Run("Update with the same record", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.NewToken()
			rec := confirm.Record{
				IsConfirmed: false,
				Token:       token,
				UserID:      user.NewID(),
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			err := repo.Upsert(ctx, rec)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			found, err := repo.FindByUserID(ctx, rec.UserID)
			if err != nil {
				t.Errorf("FindByUserID failed: %v", err)
			}
			if diff := cmp.Diff(rec, found); diff != "" {
				t.Errorf("FindByUserID returned incorrect record; (-want +got)\n%s", diff)
			}
		})

		t.Run("Update an existing record with a different token", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.NewToken()
			rec := confirm.Record{
				Token:       token,
				IsConfirmed: false,
				UserID:      user.NewID(),
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			newToken := confirm.NewToken()
			up := confirm.Record{
				IsConfirmed: rec.IsConfirmed,
				Token:       newToken,
				UserID:      rec.UserID,
			}

			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect record; (-want +got)\n%s", diff)
				}
			}

			// Find the record by token after the update.
			{
				found, err := repo.FindByToken(ctx, up.Token)
				if err != nil {
					t.Errorf("FindByToken failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByToken returned incorrect record; (-want +got)\n%s", diff)
				}
			}

			// Find the record by the _old_ token after the update.
			{
				_, err := repo.FindByToken(ctx, rec.Token)
				if err != confirm.ErrNotFound {
					t.Error("Updated record found by the old token")
				}
			}
		})

		t.Run("Update an existing record to be confirmed", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.NewToken()
			rec := confirm.Record{
				IsConfirmed: false,
				Token:       token,
				UserID:      user.NewID(),
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			up := confirm.Record{
				Token:       confirm.Token{},
				IsConfirmed: true,
				UserID:      rec.UserID,
			}

			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByUserID returned incorrect record; (-want +got)\n%s", diff)
				}
			}

			// Find the record by the _old_ token after the update.
			{
				_, err := repo.FindByToken(ctx, rec.Token)
				if err != confirm.ErrNotFound {
					t.Error("Updated record found by the old token")
				}
			}
		})
	})
}

func mustRepositoryUpsert(
	t *testing.T,
	ctx context.Context,
	repo confirm.Repository,
	rec confirm.Record,
) {
	t.Helper()
	if err := repo.Upsert(ctx, rec); err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}
}

func mustRepositoryUpsertMany(
	t *testing.T,
	ctx context.Context,
	repo confirm.Repository,
	records []confirm.Record,
) {
	t.Helper()
	for i, rec := range records {
		if err := repo.Upsert(ctx, rec); err != nil {
			t.Fatalf("records[%d]: Upsert failed: %v", i, err)
		}
	}
}
