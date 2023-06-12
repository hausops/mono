package local_test

import (
	"context"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

func TestConfirmRepository(t *testing.T) {
	ctx := context.Background()

	t.Run("FindByEmail", func(t *testing.T) {
		repo := local.NewConfirmRepository()

		records := generateTestRecords(t, 3)
		for i, rec := range records {
			if err := repo.Upsert(ctx, rec); err != nil {
				t.Fatalf("Upsert records[%d] failed: %v", i, err)
			}
		}

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != confirm.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			found, err := repo.FindByEmail(ctx, rec.Email)
			if err != nil {
				t.Errorf("FindByEmail records[%d] failed: %v", i, err)
			}
			if found != rec {
				t.Errorf("FindByEmail records[%d] returned incorrect record", i)
			}
		}
	})

	t.Run("FindByToken", func(t *testing.T) {
		repo := local.NewConfirmRepository()

		records := generateTestRecords(t, 3)
		for i, rec := range records {
			if err := repo.Upsert(ctx, rec); err != nil {
				t.Fatalf("Upsert record[%d] failed: %v", i, err)
			}
		}

		_, err := repo.FindByToken(ctx, confirm.GenerateToken())
		if err != confirm.ErrNotFound {
			t.Errorf("FindByToken(randomToken) error = %v, want: ErrNotFound", err)
		}

		for i, rec := range records {
			if rec.Token == nil {
				continue
			}
			found, err := repo.FindByToken(ctx, *rec.Token)
			if err != nil {
				t.Errorf("FindByToken records[%d] failed: %v", i, err)
			}
			if found != rec {
				t.Errorf("FindByToken records[%d] returned incorrect record", i)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new record", func(t *testing.T) {
			repo := local.NewConfirmRepository()

			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       &token,
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
				if found != rec {
					t.Error("FindByEmail returned incorrect record after insert")
				}
			}

			// Find the record by token after the insert.
			{
				found, err := repo.FindByToken(ctx, token)
				if err != nil {
					t.Errorf("FindByToken failed: %v", err)
				}
				if found != rec {
					t.Error("FindByToken returned incorrect record after insert")
				}
			}
		})

		t.Run("Update with the same record", func(t *testing.T) {
			repo := local.NewConfirmRepository()

			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       &token,
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
			if found != rec {
				t.Error("FindByEmail returned incorrect record")
			}
		})

		t.Run("Update an existing record with a different token", func(t *testing.T) {
			repo := local.NewConfirmRepository()

			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       &token,
				IsConfirmed: false,
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			newToken := confirm.GenerateToken()
			up := confirm.Record{
				Email:       rec.Email,
				Token:       &newToken,
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
				if found != up {
					t.Error("FindByEmail returned incorrect record after update")
				}
			}

			// Find the record by token after the update.
			{
				found, err := repo.FindByToken(ctx, *up.Token)
				if err != nil {
					t.Errorf("FindByToken failed: %v", err)
				}
				if found != up {
					t.Error("FindByToken returned incorrect record after update")
				}
			}

			// Find the record by the _old_ token after the update.
			{
				_, err := repo.FindByToken(ctx, *rec.Token)
				if err != confirm.ErrNotFound {
					t.Error("Updated record found by the old token")
				}
			}
		})

		t.Run("Update an existing record to be confirmed", func(t *testing.T) {
			repo := local.NewConfirmRepository()

			token := confirm.GenerateToken()
			rec := confirm.Record{
				Email:       mail.Address{Address: gofakeit.Email()},
				Token:       &token,
				IsConfirmed: false,
			}
			mustRepositoryUpsert(t, ctx, repo, rec)

			up := confirm.Record{
				Email:       rec.Email,
				Token:       nil,
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
				if found != up {
					t.Error("FindByEmail returned incorrect record after update")
				}
			}

			// Find the record by the _old_ token after the update.
			{
				_, err := repo.FindByToken(ctx, *rec.Token)
				if err != confirm.ErrNotFound {
					t.Error("Updated record found by the old token")
				}
			}
		})
	})
}

func generateTestRecords(t *testing.T, count int) []confirm.Record {
	t.Helper()
	records := make([]confirm.Record, count)
	for i := 0; i < len(records); i++ {
		token := confirm.GenerateToken()
		records[i] = confirm.Record{
			Email:       mail.Address{Address: gofakeit.Email()},
			Token:       &token,
			IsConfirmed: false,
		}
	}
	return records
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
