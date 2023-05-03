package property_test

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

// TODO: test MultiFamilyProperty
func TestPropertyService(t *testing.T) {
	t.Parallel()

	t.Run("Create", func(t *testing.T) {
		svc := newTestPropertyService(t)

		input := newFakeSingleFamilyPropertyInput(t)
		created, err := svc.Create(context.Background(), input)
		if err != nil {
			t.Errorf("Create(...) = %q; want no error", err)
		}

		p := created.(property.SingleFamilyProperty)

		if p.ID == (uuid.UUID{}) {
			t.Error("p.ID is empty")
		}

		if p.DateCreated.IsZero() {
			t.Error("p.DateCreated is empty")
		}

		if p.DateUpdated.IsZero() {
			t.Error("p.DateUpdated is empty")
		}

		if p.Unit.ID == (uuid.UUID{}) {
			t.Error("p.Unit.ID is empty")
		}

		if p.Unit.DateCreated.IsZero() {
			t.Error("p.Unit.DateCreated is empty")
		}

		if p.Unit.DateUpdated.IsZero() {
			t.Error("p.Unit.DateUpdated is empty")
		}

		if diff := cmp.Diff(input, p, ignoreGenerated); diff != "" {
			t.Errorf("Create(...): (-want +got)\n%s", diff)
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		svc := newTestPropertyService(t)

		input := newFakeSingleFamilyPropertyInput(t)
		created, err := svc.Create(context.Background(), input)
		if err != nil {
			t.Fatalf("Create(...) = %q; want no error", err)
		}

		t.Run("not found", func(t *testing.T) {
			id := uuid.New()
			_, err := svc.FindByID(context.Background(), id)
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("FindByID(%s) = %q; want error %q",
					id, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			id := created.GetID()
			got, err := svc.FindByID(context.Background(), id)
			if err != nil {
				t.Errorf("FindByID(%s) = %q; want no error", id, err)
			}
			if got != created {
				t.Errorf("FindByID(%s) = %v; want %v", id, got, created)
			}
		})
	})

	t.Run("List", func(t *testing.T) {
		svc := newTestPropertyService(t)

		var createdProperties []property.Property
		for _, input := range []property.Property{
			newFakeSingleFamilyPropertyInput(t),
			newFakeSingleFamilyPropertyInput(t),
			newFakeSingleFamilyPropertyInput(t),
		} {
			created, err := svc.Create(context.Background(), input)
			if err != nil {
				t.Fatalf("Create(...) = %q; want no error", err)
			}
			createdProperties = append(createdProperties, created)
		}

		got, err := svc.List(context.Background())
		if err != nil {
			t.Errorf("List() = %q; want no error", err)
		}

		if diff := cmp.Diff(createdProperties, got); diff != "" {
			t.Errorf("List(): (-want +got)\n%s", diff)
		}
	})

	t.Run("Update", func(t *testing.T) {
		svc := newTestPropertyService(t)

		input := newFakeSingleFamilyPropertyInput(t)
		created, err := svc.Create(context.Background(), input)
		if err != nil {
			t.Fatalf("Create(...) = %q; want no error", err)
		}

		updateYearBuilt := int32(1942)
		updateUnitSize := truncate(gofakeit.Float32Range(840, 1200))
		up := property.UpdateSingleFamilyProperty{
			Address: &property.UpdateAddress{
				Line1: &gofakeit.Address().Street,
				Line2: &gofakeit.Address().Street,
				Zip:   &gofakeit.Address().Zip,
			},
			YearBuilt: &updateYearBuilt,
			Unit: &property.UpdateRentalUnit{
				Size: &updateUnitSize,
			},
		}

		t.Run("not found", func(t *testing.T) {
			id := uuid.New()
			_, err := svc.Update(context.Background(), id, up)
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("FindByID(%s) = %q; want error %q",
					id, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			updated, err := svc.Update(context.Background(), created.GetID(), up)
			if err != nil {
				t.Errorf("Update(...) = %q; want no error", err)
			}

			c := created.(property.SingleFamilyProperty)
			u := updated.(property.SingleFamilyProperty)
			if c.DateUpdated == u.DateUpdated {
				t.Error("DateUpdated should change.")
			}

			want := c
			want.Address.Line1 = *up.Address.Line1
			want.Address.Line2 = *up.Address.Line2
			want.Address.Zip = *up.Address.Zip
			want.YearBuilt = *up.YearBuilt
			want.Unit.Size = *up.Unit.Size

			if diff := cmp.Diff(want, updated, ignoreGenerated); diff != "" {
				t.Errorf("Update(...): (-want +got)\n%s", diff)
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		p := newFakeSingleFamilyPropertyInput(t)
		repo := local.
			NewPropertyRepository().
			ReplaceProperties([]property.Property{p})

		svc := property.NewService(repo)

		t.Run("not found", func(t *testing.T) {
			_, err := svc.Delete(context.Background(), uuid.New())
			if !errors.Is(err, property.ErrNotFound) {
				t.Errorf("Delete(%s) = %q; want error %q",
					p.ID, err, property.ErrNotFound)
			}
		})

		t.Run("found", func(t *testing.T) {
			got, err := svc.Delete(context.Background(), p.ID)
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
				t.Errorf("FindByID(%s) = %q; want %q", p.ID, err, property.ErrNotFound)
			}
		})
	})
}

func newTestPropertyService(t *testing.T) *property.Service {
	t.Helper()
	return property.NewService(local.NewPropertyRepository())
}

func newFakeSingleFamilyPropertyInput(t *testing.T) property.SingleFamilyProperty {
	t.Helper()
	return property.SingleFamilyProperty{
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
			Bedrooms:  truncate(gofakeit.Float32Range(0, 3)),
			Bathrooms: truncate(gofakeit.Float32Range(0, 3)),
			Size:      truncate(gofakeit.Float32Range(320, 840)),
		},
	}
}

// cmp.Option to ignore generated fields when comparing the output
// using cmp.Diff()
var ignoreGenerated = cmp.FilterPath(func(p cmp.Path) bool {
	switch p.String() {
	case "ID", "DateCreated", "DateUpdated",
		"Unit.ID", "Unit.DateCreated", "Unit.DateUpdated":
		return true
	}
	return false
}, cmp.Ignore())

func truncate(v float32) float32 {
	return float32(math.Floor(float64(v)))
}
