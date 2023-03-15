package local

import (
	"github.com/google/uuid"
	"github.com/hausops/mono/apps/dashboard-api/domain/property"
)

type PropertyService struct {
	byId map[string]property.Property
}

func NewPropertyService() *PropertyService {
	return &PropertyService{byId: make(map[string]property.Property)}
}

func (r *PropertyService) CreateSingleFamilyProperty(in property.CreateSingleFamilyPropertyInput) (*property.SingleFamilyProperty, error) {
	unit := property.SingleFamilyPropertyUnit{
		ID:         uuid.New().String(),
		Bedrooms:   in.Unit.Bedrooms,
		Bathrooms:  in.Unit.Bathrooms,
		Size:       in.Unit.Size,
		RentAmount: in.Unit.RentAmount,
	}

	p := property.SingleFamilyProperty{
		ID:            uuid.New().String(),
		CoverImageURL: in.CoverImageURL,
		Address:       property.Address(in.Address),
		BuildYear:     in.BuildYear,
		Unit:          unit,
	}
	r.byId[p.ID] = p
	return &p, nil
}

func (r *PropertyService) CreateMultiFamilyProperty(in property.CreateMultiFamilyPropertyInput) (*property.MultiFamilyProperty, error) {
	units := make([]property.MultiFamilyPropertyUnit, 0, len(in.Units))
	for _, iu := range in.Units {
		u := property.MultiFamilyPropertyUnit{
			ID:         uuid.New().String(),
			Number:     iu.Number,
			Bedrooms:   iu.Bedrooms,
			Bathrooms:  iu.Bathrooms,
			Size:       iu.Size,
			RentAmount: iu.RentAmount,
		}
		units = append(units, u)
	}

	p := property.MultiFamilyProperty{
		ID:            uuid.New().String(),
		CoverImageURL: in.CoverImageURL,
		Address:       property.Address(in.Address),
		BuildYear:     in.BuildYear,
		Units:         units,
	}
	r.byId[p.ID] = p
	return &p, nil
}

func (r *PropertyService) FindAll() ([]property.Property, error) {
	ps := make([]property.Property, 0, len(r.byId))
	for _, p := range r.byId {
		ps = append(ps, p)
	}
	return ps, nil
}

var _ property.Service = (*PropertyService)(nil)
