// Package testing provide test clients and helpers for the property domain.
package testing

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

func newTestSingleFamilyProperty(t *testing.T) property.SingleFamilyProperty {
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
