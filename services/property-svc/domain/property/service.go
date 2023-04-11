package property

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

// Service contains the core logic. It exposes a set of public CRUD APIs.
type Service struct {
	repo Repository
}

// NewService creates a property.Service with dependencies.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create adds a Property to the Repository. It returns the created Property
// with fields like ID and DateCreated populated.
func (s *Service) Create(ctx context.Context, p Property) (Property, error) {
	now := time.Now().UTC()
	switch t := p.(type) {
	case SingleFamilyProperty:
		t.ID = uuid.New()
		t.DateCreated = now
		t.Unit.ID = uuid.New()
		t.Unit.DateCreated = now
		return s.repo.Upsert(ctx, t)
	case MultiFamilyProperty:
		t.ID = uuid.New()
		t.DateCreated = now
		for i := range t.Units {
			t.Units[i].ID = uuid.New()
			t.Units[i].DateCreated = now
		}
		return s.repo.Upsert(ctx, t)
	default:
		return nil, &UnhandledPropertyTypeError{Property: t}
	}
}

// FindByID returns the Property identified by a given ID from the Repository.
func (s *Service) FindByID(ctx context.Context, id string) (Property, error) {
	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	return s.repo.FindByID(ctx, pid)
}

// List returns all Properties from the Repository.
func (s *Service) List(ctx context.Context) ([]Property, error) {
	return s.repo.List(ctx)
}

// Update merges in to the saved property matching id and saves it back
// to the Repository.
func (s *Service) Update(ctx context.Context, id string, up UpdateProperty) (Property, error) {
	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	saved, err := s.repo.FindByID(ctx, pid)
	if err != nil {
		return nil, ErrNotFound
	}

	now := time.Now().UTC()
	switch t := saved.(type) {
	case SingleFamilyProperty:
		if err := mapstructure.Decode(up, &t); err != nil {
			return nil, err
		}
		t.DateUpdated = now
		return s.repo.Upsert(ctx, t)
	case MultiFamilyProperty:
		if err := mapstructure.Decode(up, &t); err != nil {
			return nil, err
		}
		t.DateUpdated = now
		return s.repo.Upsert(ctx, t)
	default:
		return nil, &UnhandledPropertyTypeError{Property: t}
	}

	// switch t := in.(type) {
	// case SingleFamilyProperty:
	// 	sp, ok := saved.(SingleFamilyProperty)
	// 	if !ok {
	// 		return nil, fmt.Errorf("property type mismatch: saved[%T], request[%T]", sp, t)
	// 	}
	// 	if err := mapstructure.Decode(t, &sp); err != nil {
	// 		return nil, err
	// 	}
	// 	sp.DateUpdated = now
	// 	sp.Unit.DateUpdated = now
	// 	return s.repo.Upsert(ctx, sp)
	// case MultiFamilyProperty:
	// 	mp, ok := saved.(MultiFamilyProperty)
	// 	if !ok {
	// 		return nil, fmt.Errorf("property type mismatch: saved[%T], request[%T]", mp, t)
	// 	}
	// 	if err := mapstructure.Decode(t, &mp); err != nil {
	// 		return nil, err
	// 	}
	// 	mp.DateUpdated = now
	// 	for i := range mp.Units {
	// 		mp.Units[i].DateUpdated = now
	// 	}
	// 	return s.repo.Upsert(ctx, mp)
	// default:
	// 	return nil, &UnhandledPropertyTypeError{Property: t}
	// }
}

// Delete removes the Property identified by a given ID from the Repository.
func (s *Service) Delete(ctx context.Context, id string) (Property, error) {
	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return s.repo.Delete(ctx, pid)
}
