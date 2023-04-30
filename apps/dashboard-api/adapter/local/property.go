package local

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hausops/mono/apps/dashboard-api/domain/property"
	"github.com/mitchellh/mapstructure"
)

type PropertyService struct {
	byId map[string]property.Property
}

var _ property.Service = (*PropertyService)(nil)

func NewPropertyService() *PropertyService {
	return &PropertyService{byId: make(map[string]property.Property)}
}

func (r *PropertyService) CreateSingleFamilyProperty(
	_ context.Context,
	in property.CreateSingleFamilyPropertyInput,
) (*property.SingleFamilyProperty, error) {
	unit := property.SingleFamilyPropertyUnit{
		ID:         uuid.New().String(),
		Bedrooms:   in.Unit.Bedrooms,
		Bathrooms:  in.Unit.Bathrooms,
		Size:       in.Unit.Size,
		RentAmount: in.Unit.RentAmount,
	}

	p := &property.SingleFamilyProperty{
		ID:            uuid.New().String(),
		CoverImageURL: in.CoverImageURL,
		Address:       property.Address(in.Address),
		YearBuilt:     in.YearBuilt,
		Unit:          unit,
	}
	r.byId[p.ID] = p
	return p, nil
}

func (r *PropertyService) CreateMultiFamilyProperty(
	_ context.Context,
	in property.CreateMultiFamilyPropertyInput,
) (*property.MultiFamilyProperty, error) {
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

	p := &property.MultiFamilyProperty{
		ID:            uuid.New().String(),
		CoverImageURL: in.CoverImageURL,
		Address:       property.Address(in.Address),
		YearBuilt:     in.YearBuilt,
		Units:         units,
	}
	r.byId[p.ID] = p
	return p, nil
}

func (r *PropertyService) FindByID(_ context.Context, id string) (property.Property, error) {
	p, ok := r.byId[id]
	if !ok {
		return nil, property.NotFoundError{ID: id}
	}
	return p, nil
}

func (r *PropertyService) FindAll(_ context.Context) ([]property.Property, error) {
	ps := make([]property.Property, 0, len(r.byId))
	for _, p := range r.byId {
		ps = append(ps, p)
	}
	return ps, nil
}

func (r *PropertyService) UpdateSingleFamilyPropertyByID(
	_ context.Context,
	id string,
	in property.UpdateSingleFamilyPropertyInput,
) (*property.SingleFamilyProperty, error) {
	p, ok := r.byId[id]
	if !ok {
		return nil, property.NotFoundError{ID: id}
	}
	sp, ok := p.(*property.SingleFamilyProperty)
	if !ok {
		return nil, errors.New("update property (single-family): property type mismatch")
	}
	if err := mapstructure.Decode(in, &sp); err != nil {
		return nil, err
	}
	r.byId[id] = sp
	return sp, nil
}

func (r *PropertyService) UpdateMultiFamilyPropertyByID(
	_ context.Context,
	id string,
	in property.UpdateMultiFamilyPropertyInput,
) (*property.MultiFamilyProperty, error) {
	p, ok := r.byId[id]
	if !ok {
		return nil, property.NotFoundError{ID: id}
	}
	mp, ok := p.(*property.MultiFamilyProperty)
	if !ok {
		return nil, errors.New("update property (multi-family): property type mismatch")
	}
	if err := mapstructure.Decode(in, &mp); err != nil {
		return nil, err
	}
	r.byId[id] = mp
	return mp, nil
}

func (r *PropertyService) DeleteByID(_ context.Context, id string) (property.Property, error) {
	p, ok := r.byId[id]
	if !ok {
		return nil, property.NotFoundError{ID: id}
	}
	delete(r.byId, id)
	return p, nil
}
