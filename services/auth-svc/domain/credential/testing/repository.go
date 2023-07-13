package testing

import (
	"bytes"
	"context"
	"net/mail"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
)

// TestRepository is a suite of unit tests that ensure
// the implementation of credential.Repository adheres to the expected behaviors.
//
// newRepo is a factory function that should return the concrete implementation
// of credential.Repository under test.
func TestRepository(t *testing.T, newRepo func(t *testing.T) credential.Repository) {
	ctx := context.Background()

	t.Run("FindByEmail", func(t *testing.T) {
		repo := newRepo(t)
		creds := generateTestCredentials(t, 3)
		for i, cred := range creds {
			if err := repo.Upsert(ctx, cred); err != nil {
				t.Fatalf("credential[%d]:Upsert failed: %v", i, err)
			}
		}

		_, err := repo.FindByEmail(ctx, mail.Address{Address: gofakeit.Email()})
		if err != credential.ErrNotFound {
			t.Errorf("FindByEmail(randomEmail) error = %v, want: ErrNotFound", err)
		}

		for i, cred := range creds {
			found, err := repo.FindByEmail(ctx, cred.Email)
			if err != nil {
				t.Errorf("credential[%d]: FindByEmail failed: %v", i, err)
			}
			if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
				t.Errorf("credential[%d]: FindByEmail returned incorrect credential; (-want +got)\n%s", i, diff)
			}
			if !bytes.Equal(found.Password, cred.Password) {
				t.Errorf("credential[%d]: FindByEmail returned incorrect password", i)
			}
		}
	})

	t.Run("FindByUserID", func(t *testing.T) {
		repo := newRepo(t)
		creds := generateTestCredentials(t, 3)
		for i, cred := range creds {
			if err := repo.Upsert(ctx, cred); err != nil {
				t.Fatalf("credential[%d]:Upsert failed: %v", i, err)
			}
		}

		_, err := repo.FindByUserID(ctx, generateTestCredential(t).UserID)
		if err != credential.ErrNotFound {
			t.Errorf("FindByUserID(randomUserID) error = %v, want: ErrNotFound", err)
		}

		for i, cred := range creds {
			found, err := repo.FindByUserID(ctx, cred.UserID)
			if err != nil {
				t.Errorf("credential[%d]: FindByUserID failed: %v", i, err)
			}
			if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
				t.Errorf("credential[%d]: FindByUserID returned incorrect credential; (-want +got)\n%s", i, diff)
			}
			if !bytes.Equal(found.Password, cred.Password) {
				t.Errorf("credential[%d]: FindByUserID returned incorrect password", i)
			}
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("Insert a new credential", func(t *testing.T) {
			repo := newRepo(t)
			cred := generateTestCredential(t)

			err := repo.Upsert(ctx, cred)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the credential by email after the insert.
			{
				found, err := repo.FindByEmail(ctx, cred.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
					t.Errorf("FindByEmail returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, cred.Password) {
					t.Error("FindByEmail returned incorrect password")
				}
			}

			// Find the credential by user ID after the insert.
			{
				found, err := repo.FindByUserID(ctx, cred.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
					t.Errorf("FindByUserID returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, cred.Password) {
					t.Error("FindByUserID returned incorrect password")
				}
			}
		})

		t.Run("Update with the same credential", func(t *testing.T) {
			repo := newRepo(t)
			cred := generateTestCredential(t)
			mustRepositoryUpsert(t, ctx, repo, cred)

			err := repo.Upsert(ctx, cred)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the credential by email after the insert.
			{
				found, err := repo.FindByEmail(ctx, cred.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
					t.Errorf("FindByEmail returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, cred.Password) {
					t.Error("FindByEmail returned incorrect password")
				}
			}

			// Find the credential by user ID after the insert.
			{
				found, err := repo.FindByUserID(ctx, cred.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(cred, found, ignorePassword); diff != "" {
					t.Errorf("FindByUserID returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, cred.Password) {
					t.Error("FindByUserID returned incorrect password")
				}
			}
		})

		t.Run("Update an existing credential with a different password", func(t *testing.T) {
			repo := newRepo(t)
			cred := generateTestCredential(t)
			mustRepositoryUpsert(t, ctx, repo, cred)

			up := credential.Credential{
				Email:    cred.Email,
				Password: generateTestCredential(t).Password,
				UserID:   cred.UserID,
			}
			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the credential by email after the update.
			{
				found, err := repo.FindByEmail(ctx, up.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(up, found, ignorePassword); diff != "" {
					t.Errorf("FindByEmail returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, up.Password) {
					t.Error("FindByEmail returned incorrect password")
				}
			}

			// Find the credential by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found, ignorePassword); diff != "" {
					t.Errorf("FindByUserID returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, up.Password) {
					t.Error("FindByUserID returned incorrect password")
				}
			}
		})

		t.Run("Update an existing credential with a different email", func(t *testing.T) {
			repo := newRepo(t)
			cred := generateTestCredential(t)
			mustRepositoryUpsert(t, ctx, repo, cred)

			up := credential.Credential{
				Email:    generateTestCredential(t).Email,
				Password: cred.Password,
				UserID:   cred.UserID,
			}
			err := repo.Upsert(ctx, up)
			if err != nil {
				t.Errorf("Upsert failed: %v", err)
			}

			// Find the credential by email after the update.
			{
				found, err := repo.FindByEmail(ctx, up.Email)
				if err != nil {
					t.Errorf("FindByEmail failed: %v", err)
				}
				if diff := cmp.Diff(up, found, ignorePassword); diff != "" {
					t.Errorf("FindByEmail returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, up.Password) {
					t.Error("FindByEmail returned incorrect password")
				}
			}

			// Find the credential by user ID after the update.
			{
				found, err := repo.FindByUserID(ctx, up.UserID)
				if err != nil {
					t.Errorf("FindByUserID failed: %v", err)
				}
				if diff := cmp.Diff(up, found, ignorePassword); diff != "" {
					t.Errorf("FindByUserID returned incorrect credential; (-want +got)\n%s", diff)
				}
				if !bytes.Equal(found.Password, up.Password) {
					t.Error("FindByUserID returned incorrect password")
				}
			}

			// Find the credential by the _old_ email after the update.
			{
				_, err := repo.FindByEmail(ctx, cred.Email)
				if err != credential.ErrNotFound {
					t.Error("Updated session found by the old user ID")
				}
			}
		})
	})
}

// cmp.Option to ignore the Password field when comparing the output
// using cmp.Diff()
var ignorePassword = cmp.FilterPath(func(p cmp.Path) bool {
	return p.String() == "Password"
}, cmp.Ignore())

func mustRepositoryUpsert(
	t *testing.T,
	ctx context.Context,
	repo credential.Repository,
	cred credential.Credential,
) {
	t.Helper()
	if err := repo.Upsert(ctx, cred); err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}
}
