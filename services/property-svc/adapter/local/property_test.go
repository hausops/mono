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
			_, err := repo.Delete(context.Background(), uuid.New())
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("Delete(%s) = %q; want error %q", p.ID, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			got, err := repo.Delete(context.Background(), p.ID)
			if err != nil {
				t.Log("On delete success, does not return an error.")
				t.Errorf("Delete(%s) = %q; want no error", p.ID, err)
			}
			if got != p {
				t.Log("On delete success, returns the deleted property.")
				t.Errorf("Delete(%s) = %v; want %v", p.ID, got, p)
			}

			_, err = repo.FindByID(context.Background(), p.ID)
			if !errors.Is(err, property.ErrNotFound) {
				t.Log("The deleted property should not longer be found.")
				t.Errorf("FindByID(%s) = %q; want error %q",
					p.ID, err, property.ErrNotFound)
			}
		})
	})

	t.Run("FindByID", func(t *testing.T) {
		p := newFakeSingleFamilyProperty(t)
		repo := local.
			NewPropertyRepository().
			ReplaceProperties([]property.Property{p})

		t.Run("not found", func(t *testing.T) {
			id := uuid.New()
			_, err := repo.FindByID(context.Background(), id)
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("FindByID(%s) = %q; want error %q",
					id, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			id := p.ID
			got, err := repo.FindByID(context.Background(), id)
			if err != nil {
				t.Errorf("FindByID(%s) = %q; want no error", id, err)
			}
			if got != p {
				t.Errorf("FindByID(%s) = %v; want %v", id, got, p)
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

		got, err := repo.List(context.Background())
		if err != nil {
			t.Errorf("List() = %q; want no error", err)
		}
		if diff := cmp.Diff(ps, got); diff != "" {
			t.Errorf("List(): (-want +got)\n%s", diff)
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("insert new properties", func(t *testing.T) {
			repo := local.NewPropertyRepository()

			for i, p := range []property.SingleFamilyProperty{
				newFakeSingleFamilyProperty(t),
				newFakeSingleFamilyProperty(t),
			} {
				got, err := repo.Upsert(context.Background(), p)
				if err != nil {
					t.Logf("[%d] On upsert success, does not return an error.", i)
					t.Errorf("Upsert(%v) = %q; want no error", p, err)
				}
				if got != p {
					t.Logf("[%d] On upsert success, returns the upserted property.", i)
					t.Errorf("Upsert(%v) = %v; want %v", p, got, p)
				}

				got, _ = repo.FindByID(context.Background(), p.ID)
				if got != p {
					t.Logf("[%d] The upserted property should be found.", i)
					t.Errorf("FindByID(%s) = %v; want %v", p.ID, got, p)
				}
			}
		})

		t.Run("replace existing properties", func(t *testing.T) {
			p1 := newFakeSingleFamilyProperty(t)
			p2 := newFakeSingleFamilyProperty(t)

			repo := local.
				NewPropertyRepository().
				ReplaceProperties([]property.Property{p1, p2})

			updateProperty := newFakeSingleFamilyProperty(t)
			updateProperty.ID = p1.ID

			got, err := repo.Upsert(context.Background(), updateProperty)
			if err != nil {
				t.Log("On upsert success, does not return an error.")
				t.Errorf("Upsert(%v) = %q; want no error", updateProperty, err)
			}
			if got != updateProperty {
				t.Log("On upsert success, returns the upserted property.")
				t.Errorf("Upsert(%v) = %v; want %v", updateProperty, got, updateProperty)
			}

			got, _ = repo.FindByID(context.Background(), updateProperty.ID)
			if got != updateProperty {
				t.Log("The upserted property should be found.")
				t.Errorf("FindByID(%s) = %v; want %v",
					updateProperty.ID, got, updateProperty)
			}
		})

		t.Run("when Upsert with nil", func(t *testing.T) {
			repo := local.NewPropertyRepository()

			_, err := repo.Upsert(context.Background(), nil)
			if err == nil {
				t.Error("Upsert(<nil>) should return an error.")
			}
		})

		t.Run("when Upsert with no property ID", func(t *testing.T) {
			p := newFakeSingleFamilyProperty(t)
			p.ID = uuid.UUID{}
			repo := local.NewPropertyRepository()

			_, err := repo.Upsert(context.Background(), p)
			if err == nil {
				t.Error("Upsert(...) should return an error.")
			}
		})
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
