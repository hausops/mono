package local

import (
	"github.com/google/uuid"
	"github.com/hausops/mono/apps/dashboard-api/domain/property"
)

type PropertyRepository struct {
	byId map[string]property.Property
}

func NewPropertyRepository() *PropertyRepository {
	return &PropertyRepository{byId: make(map[string]property.Property)}
}

func (r *PropertyRepository) CreateSingleFamilyProperty(input property.CreateSingleFamilyPropertyInput) (*property.SingleFamilyProperty, error) {
	p := input.ToProperty()
	p.ID = uuid.New().String()
	p.Unit.ID = uuid.New().String()
	r.byId[p.ID] = p
	return &p, nil
}

func (r *PropertyRepository) CreateMultiFamilyProperty(input property.CreateMultiFamilyPropertyInput) (*property.MultiFamilyProperty, error) {
	var units []property.MultiFamilyPropertyUnit
	for _, iu := range input.Units {
		u := iu.ToUnit()
		u.ID = uuid.New().String()
		units = append(units, u)
	}
	p := input.ToProperty()
	p.ID = uuid.New().String()
	p.Units = units
	r.byId[p.ID] = p
	return &p, nil
}

func (r *PropertyRepository) FindAll() ([]property.Property, error) {
	var ps []property.Property
	for _, p := range r.byId {
		ps = append(ps, p)
	}
	return ps, nil
}

var _ property.Repository = (*PropertyRepository)(nil)
