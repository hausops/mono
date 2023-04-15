package local_test

import (
	"context"
	"errors"
	"math"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

// TODO: move this to a suite of contract tests exported by domain (property.Repository)
// so concrete implementations can run the suite to ensure the implementation
// conforms to the expected behavior.
func TestPropertyRepository(t *testing.T) {
	t.Parallel()

	t.Run("Delete", func(t *testing.T) {
		p := newFakeSingleFamilyProperty(t)
		repo := local.
			NewPropertyRepository().
			ReplaceProperties([]property.Property{p})

		t.Run("not found", func(t *testing.T) {
			_, err := repo.Delete(context.TODO(), uuid.New())
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("Delete(%s) = %q; want %q", p.ID, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			t.Log("On delete success, returns the deleted property.")
			got, err := repo.Delete(context.TODO(), p.ID)
			if err != nil {
				t.Errorf("Delete(%s) = %q; want no error", p.ID, err)
			}
			if got != p {
				t.Errorf("Delete(%s) = %v; want %v", p.ID, got, p)
			}

			t.Log("The deleted property should not longer be found in the repo.")
			_, err = repo.FindByID(context.TODO(), p.ID)
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("...(%s) = %q; want %q", p.ID, err, property.ErrNotFound)
			}
		})
	})

	t.Run("FindByID", func(t *testing.T) {
		p := newFakeSingleFamilyProperty(t)
		repo := local.
			NewPropertyRepository().
			ReplaceProperties([]property.Property{p})

		t.Run("not found", func(t *testing.T) {
			_, err := repo.FindByID(context.TODO(), uuid.New())
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("FindByID(%s) = %q; want %q", p.ID, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			got, err := repo.FindByID(context.TODO(), p.ID)
			if err != nil {
				t.Errorf("FindByID(%s) = %q; want no error", p.ID, err)
			}
			if got != p {
				t.Errorf("FindByID(%s) = %v; want %v", p.ID, got, p)
			}
		})
	})

	t.Run("List", func(t *testing.T) {
		ps := []property.Property{
			newFakeSingleFamilyProperty(t),
			newFakeSingleFamilyProperty(t),
			newFakeSingleFamilyProperty(t),
		}

		repo := local.
			NewPropertyRepository().
			ReplaceProperties(ps)

		got, err := repo.List(context.TODO())
		if err != nil {
			t.Errorf("List() = %q; want no error", err)
		}
		if diff := cmp.Diff(ps, got); diff != "" {
			t.Errorf("List(): (-want +got)\n%s", diff)
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		ps := []property.SingleFamilyProperty{
			newFakeSingleFamilyProperty(t),
			newFakeSingleFamilyProperty(t),
		}

		repo := local.NewPropertyRepository()

		for i, p := range ps {
			t.Log("On upsert success, returns the upserted property.")
			got, err := repo.Upsert(context.TODO(), p)
			if err != nil {
				t.Errorf("[%d] Upsert(%v) = %q; want no error", i, p, err)
			}
			if got != p {
				t.Errorf("[%d] Upsert(%v) = %v; want %v", i, p, got, p)
			}
		}
	})
}

func newFakeSingleFamilyProperty(t *testing.T) property.SingleFamilyProperty {
	t.Helper()
	return property.SingleFamilyProperty{
		ID: uuid.New(),
		Address: property.Address{
			Line1: gofakeit.Street(),
			City:  gofakeit.City(),
			State: gofakeit.StateAbr(),
			Zip:   gofakeit.Zip(),
		},
		CoverImageURL: "https://hausops.com/images/example-sfp.jpg",
		// it is okay if the values don't make much sense in real life
		// we only need the correct data types for testing
		Unit: property.RentalUnit{
			ID:        uuid.New(),
			Bedrooms:  truncate(gofakeit.Float32Range(0, 3)),
			Bathrooms: truncate(gofakeit.Float32Range(0, 3)),
			Size:      truncate(gofakeit.Float32Range(320, 840)),
		},
	}
}

func truncate(v float32) float32 {
	return float32(math.Floor(float64(v)))
}
