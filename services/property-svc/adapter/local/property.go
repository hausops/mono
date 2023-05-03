package local

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

type propertyRepository struct {
	byID map[uuid.UUID]property.Property
}

// NewPropertyRepository creates a new instance of the property repository with
// an initial empty state.
//
// The returned property repository can be used to store and retrieve properties.
func NewPropertyRepository() *propertyRepository {
	return &propertyRepository{
		byID: make(map[uuid.UUID]property.Property),
	}
}

// Ensure propertyRepository implements the property.Repository interface.
var _ property.Repository = (*propertyRepository)(nil)

func (r *propertyRepository) Delete(_ context.Context, id uuid.UUID) (property.Property, error) {
	p, ok := r.byID[id]
	if !ok {
		return nil, property.ErrNotFound
	}
	delete(r.byID, id)
	return p, nil
}

func (r *propertyRepository) FindByID(_ context.Context, id uuid.UUID) (property.Property, error) {
	p, ok := r.byID[id]
	if !ok {
		return nil, property.ErrNotFound
	}
	return p, nil
}

func (r *propertyRepository) List(_ context.Context) ([]property.Property, error) {
	ps := make([]property.Property, 0, len(r.byID))
	for _, p := range r.byID {
		ps = append(ps, p)
	}
	return ps, nil
}

func (r *propertyRepository) Upsert(_ context.Context, p property.Property) (property.Property, error) {
	if p == nil {
		return nil, fmt.Errorf("invalid parameter: property is nil")
	}

	if p.GetID() == (uuid.UUID{}) {
		return nil, property.MissingIDError{Message: "missing property ID"}
	}

	switch t := p.(type) {
	case property.SingleFamilyProperty:
		if t.Unit.ID == (uuid.UUID{}) {
			return nil, property.MissingIDError{Message: "missing unit ID"}
		}
	case property.MultiFamilyProperty:
		for _, unit := range t.Units {
			if unit.ID == (uuid.UUID{}) {
				return nil, property.MissingIDError{Message: "missing unit ID"}
			}
		}
	}

	r.byID[p.GetID()] = p
	return p, nil
}

/**
 * Utility methods specific this implementation.
 */

// ReplaceProperties replaces all properties in the repository with
// the properties in ps.
//
// This method is intended to be used for populating the repository
// with initial entries.
func (r *propertyRepository) ReplaceProperties(ps []property.Property) *propertyRepository {
	r.byID = make(map[uuid.UUID]property.Property, len(ps))
	for _, p := range ps {
		r.byID[p.GetID()] = p
	}
	return r
}

// ExampleProperties returns an example slice of property.Property.
//
// This function is intended to be used for testing and demonstration purposes.
func ExampleProperties() []property.Property {
	return []property.Property{
		property.SingleFamilyProperty{
			ID: uuid.MustParse("7f9dbb2e-fde0-4ea8-b21a-1236960bda59"),
			Address: property.Address{
				Line1: "527 Bridle Street",
				City:  "Flowery Branch",
				State: "GA",
				Zip:   "30542",
			},
			CoverImageURL: "/images/pexels-scott-webb-1029599.jpg",
			Unit: property.RentalUnit{
				ID:        uuid.MustParse("5e3f8f3c-763a-453f-8208-66a45b47c6af"),
				Bedrooms:  3,
				Bathrooms: 2.5,
				Size:      1024,
			},
		},

		property.SingleFamilyProperty{
			ID: uuid.MustParse("1f179de1-9089-4cbb-a74e-44181a244c3b"),
			Address: property.Address{
				Line1: "495 Ohio Street",
				City:  "Harleysville",
				State: "PA",
				Zip:   "19438",
			},
			CoverImageURL: "/images/pexels-mark-mccammon-2724749.jpg",
			Unit: property.RentalUnit{
				ID: uuid.MustParse("efa7f295-1c07-4fa7-81af-81a2ab8e8626"),
			},
		},

		property.SingleFamilyProperty{
			ID: uuid.MustParse("eed9a5c4-7aad-4373-8e45-5ee97ccb83e3"),
			Address: property.Address{
				Line1: "9026 Washington Dr.",
				City:  "Orland Park",
				State: "IL",
				Zip:   "60462",
			},
			CoverImageURL: "/images/pexels-curtis-adams-3288102.jpg",
			Unit: property.RentalUnit{
				ID: uuid.MustParse("b0594793-a611-4e68-8a1d-7d365e3f7f4e"),
			},
		},

		property.MultiFamilyProperty{
			ID: uuid.MustParse("425f2fc6-2d4a-4577-a194-9c81a78d405f"),
			Address: property.Address{
				Line1: "10 Rosa Street",
				City:  "San Francisco",
				State: "CA",
				Zip:   "94107",
			},
			CoverImageURL: "/images/pexels-quintin-gellar-612949.jpg",
			Units: []property.RentalUnit{
				{
					ID:         uuid.MustParse("6504d8b0-96b0-4470-8337-24a7add45915"),
					Number:     "201",
					Bedrooms:   0,
					Bathrooms:  1,
					Size:       524,
					RentAmount: 2075,
				},
				{
					ID:         uuid.MustParse("4d7b9a25-1cd8-4d06-aaf9-39abbc32a41a"),
					Number:     "301",
					Bedrooms:   2,
					Bathrooms:  2,
					Size:       950,
					RentAmount: 3850,
				},
				{
					ID:         uuid.MustParse("72d6f174-da28-4ef6-aa52-a2950b9b0bb2"),
					Number:     "302",
					Bedrooms:   2,
					Bathrooms:  2,
					Size:       982,
					RentAmount: 4000,
				},
				{
					ID:         uuid.MustParse("e62c2e3a-cec4-4373-8593-4d62b5f53db9"),
					Number:     "303",
					Bedrooms:   2,
					Bathrooms:  2,
					Size:       982,
					RentAmount: 4000,
				},
			},
		},

		property.SingleFamilyProperty{
			ID: uuid.MustParse("88d8de72-d25e-4f43-a9de-7d9c34398aa7"),
			Address: property.Address{
				Line1: "9189 South Argyle Dr.",
				City:  "Natchez",
				State: "MS",
				Zip:   "39120",
			},
			Unit: property.RentalUnit{
				ID: uuid.MustParse("573742cf-618b-4ade-b04f-6a225b6f4d2e"),
			},
		},

		property.SingleFamilyProperty{
			ID: uuid.MustParse("2534f3df-dbde-4945-bade-7901c00a3e9a"),
			Address: property.Address{
				Line1: "9190 South Argyle Dr.",
				City:  "Natchez",
				State: "MS",
				Zip:   "39120",
			},
			Unit: property.RentalUnit{
				ID: uuid.MustParse("4703522d-6f83-4471-a620-bfe4b75243f0"),
			},
		},

		property.SingleFamilyProperty{
			ID: uuid.MustParse("9e2d7930-18c9-4b57-8743-cf287118f106"),
			Address: property.Address{
				Line1: "290 County Rd",
				Line2: "#2011",
				City:  "Vista",
				State: "CA",
				Zip:   "92081",
			},
			Unit: property.RentalUnit{
				ID: uuid.MustParse("207a6e67-0f27-4802-9d7d-e257c2079ca6"),
			},
		},
	}
}
