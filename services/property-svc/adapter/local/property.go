package local

import (
	"context"

	"github.com/hausops/mono/services/property-svc/domain/property"
)

type propertyRepository struct {
	byId map[string]property.Property
}

func NewPropertyRepository() *propertyRepository {
	return &propertyRepository{
		byId: map[string]property.Property{
			"1029599": property.SingleFamilyProperty{
				ID: "1029599",
				Address: property.Address{
					Line1: "527 Bridle Street",
					City:  "Flowery Branch",
					State: "GA",
					Zip:   "30542",
				},
				CoverImageUrl: "/images/pexels-scott-webb-1029599.jpg",
				Unit: property.RentalUnit{
					ID:        "1029599-0",
					Bedrooms:  3,
					Bathrooms: 2.5,
					Size:      1024,
				},
			},
			"2724749": property.SingleFamilyProperty{
				ID: "2724749",
				Address: property.Address{
					Line1: "495 Ohio Street",
					City:  "Harleysville",
					State: "PA",
					Zip:   "19438",
				},
				CoverImageUrl: "/images/pexels-mark-mccammon-2724749.jpg",
				Unit: property.RentalUnit{
					ID: "2724749-0",
				},
			},
			"3288102": property.SingleFamilyProperty{
				ID: "3288102",
				Address: property.Address{
					Line1: "9026 Washington Dr.",
					City:  "Orland Park",
					State: "IL",
					Zip:   "60462",
				},
				CoverImageUrl: "/images/pexels-curtis-adams-3288102.jpg",
				Unit: property.RentalUnit{
					ID: "3288102-0",
				},
			},
			"4375210": property.MultiFamilyProperty{
				ID: "4375210",
				Address: property.Address{
					Line1: "10 Rosa Street",
					City:  "San Francisco",
					State: "CA",
					Zip:   "94107",
				},
				CoverImageUrl: "/images/pexels-quintin-gellar-612949.jpg",
				Units: []property.RentalUnit{
					{
						ID:         "4375210-1",
						Number:     "201",
						Bedrooms:   0,
						Bathrooms:  1,
						Size:       524,
						RentAmount: 2075,
					},
					{
						ID:         "4375210-2",
						Number:     "301",
						Bedrooms:   2,
						Bathrooms:  2,
						Size:       950,
						RentAmount: 3850,
					},
					{
						ID:         "4375210-3",
						Number:     "302",
						Bedrooms:   2,
						Bathrooms:  2,
						Size:       982,
						RentAmount: 4000,
					},
					{
						ID:         "4375210-4",
						Number:     "303",
						Bedrooms:   2,
						Bathrooms:  2,
						Size:       982,
						RentAmount: 4000,
					},
				},
			},
			"9999990": property.SingleFamilyProperty{
				ID: "9999990",
				Address: property.Address{
					Line1: "9189 South Argyle Dr.",
					City:  "Natchez",
					State: "MS",
					Zip:   "39120",
				},
				Unit: property.RentalUnit{
					ID: "9999990",
				},
			},
			"9999991": property.SingleFamilyProperty{
				ID: "9999991",
				Address: property.Address{
					Line1: "9189 South Argyle Dr.",
					City:  "Natchez",
					State: "MS",
					Zip:   "39120",
				},
				Unit: property.RentalUnit{
					ID: "9999991",
				},
			},
			"9999992": property.SingleFamilyProperty{
				ID: "9999992",
				Address: property.Address{
					Line1: "290 County Rd",
					Line2: "#2011",
					City:  "Vista",
					State: "CA",
					Zip:   "92081",
				},
				Unit: property.RentalUnit{
					ID: "9999992",
				},
			},
		},
	}
}

var _ property.Repository = (*propertyRepository)(nil)

func (r *propertyRepository) FindByID(_ context.Context, id string) (property.Property, error) {
	p, ok := r.byId[id]
	if !ok {
		return nil, property.ErrNotFound
	}
	return p, nil
}

func (r *propertyRepository) List(_ context.Context) ([]property.Property, error) {
	ps := make([]property.Property, 0, len(r.byId))
	for _, p := range r.byId {
		ps = append(ps, p)
	}
	return ps, nil
}

// func (r *propertyRepository) Upsert(p property.Property) (property.Property, error) {}
