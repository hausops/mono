package testing

import (
	"context"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of confirm.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of confirm.Repository under test.
func TestRepository(t *testing.T, newRepo func(t *testing.T) confirm.Repository) {
	ctx := context.Background()

	t.Run("FindByEmail", func(t *testing.T) {
		repo := newRepo(t)
		records := generateTestRecords(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, records)

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != confirm.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			found, err := repo.FindByEmail(ctx, rec.Email)
			if err != nil {
				t.Errorf("records[%d]: FindByEmail failed: %v", i, err)
			}
			if found != rec {
				t.Errorf("records[%d]: FindByEmail returned incorrect record", i)
			}
		}
	})

	t.Run("FindByToken", func(t *testing.T) {
		repo := newRepo(t)
		records := generateTestRecords(t, 3)
		mustRepositoryUpsertMany(t, ctx, repo, records)

		_, err := repo.FindByToken(ctx, confirm.GenerateToken())
		if err != confirm.ErrNotFound {
			t.Errorf("FindByToken(randomToken) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			found, err := repo.FindByToken(ctx, rec.Token)
			if err != nil {
				t.Errorf("records[%d]: FindByToken failed: %v", i, err)
			}
			if found != rec {
				t.Errorf("records[%d]: FindByToken returned incorrect record", i)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new record", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       token,
				IsConfirmed: false,
			}

			err := repo.Upsert(ctx, rec)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by email after the insert.
			{
				found, err := repo.FindByEmail(ctx, rec.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(rec, found); diff != "" {
					t.Errorf("FindByEmail returned incorrect record; (-want +got)\n%s", diff)
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
			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       token,
				IsConfirmed: false,
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			err := repo.Upsert(ctx, rec)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			found, err := repo.FindByEmail(ctx, rec.Email)
			if err != nil {
				t.Errorf("FindByEmail failed: %v", err)
			}
			if diff := cmp.Diff(rec, found); diff != "" {
				t.Errorf("FindByEmail returned incorrect record; (-want +got)\n%s", diff)
			}
		})

		t.Run("Update an existing record with a different token", func(t *testing.T) {
			repo := newRepo(t)
			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       token,
				IsConfirmed: false,
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			newToken := confirm.GenerateToken()
			up := confirm.Record{
				Email:       rec.Email,
				Token:       newToken,
				IsConfirmed: rec.IsConfirmed,
			}

			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by email after the update.
			{
				found, err := repo.FindByEmail(ctx, up.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByEmail returned incorrect record; (-want +got)\n%s", diff)
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
			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       token,
				IsConfirmed: false,
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			up := confirm.Record{
				Email:       rec.Email,
				Token:       confirm.Token{},
				IsConfirmed: true,
			}

			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the record by email after the update.
			{
				found, err := repo.FindByEmail(ctx, up.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(up, found); diff != "" {
					t.Errorf("FindByEmail returned incorrect record; (-want +got)\n%s", diff)
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
