package property

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

// Service contains the core logic. It exposes a set of public CRUD APIs.
type Service struct {
	repo Repository
}

// NewService creates a property.Service with its dependencies.
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
		t.DateUpdated = now
		t.Unit.ID = uuid.New()
		t.Unit.DateCreated = now
		t.Unit.DateUpdated = now
		return s.repo.Upsert(ctx, t)
	case MultiFamilyProperty:
		t.ID = uuid.New()
		t.DateCreated = now
		t.DateUpdated = now
		for i := range t.Units {
			t.Units[i].ID = uuid.New()
			t.Units[i].DateCreated = now
			t.Units[i].DateUpdated = now
		}
		return s.repo.Upsert(ctx, t)
	default:
		return nil, UnhandledPropertyTypeError{Property: t}
	}
}

// FindByID returns the Property identified by id from the Repository.
func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (Property, error) {
	return s.repo.FindByID(ctx, id)
}

// List returns all Properties from the Repository.
func (s *Service) List(ctx context.Context) ([]Property, error) {
	return s.repo.List(ctx)
}

// Update merges in to the saved Property identified by id and saves it back
// to the Repository. It returns the updated Property with fields like
// DateUpdated updated.
func (s *Service) Update(ctx context.Context, id uuid.UUID, up UpdateProperty) (Property, error) {
	saved, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	switch t := saved.(type) {
	case SingleFamilyProperty:
		if _, ok := up.(UpdateSingleFamilyProperty); !ok {
			return nil, UpdateWrongPropertyTypeError{Property: saved, UpdateProperty: up}
		}
		if err := mapstructure.Decode(up, &t); err != nil {
			return nil, err
		}
		t.DateUpdated = time.Now().UTC()
		return s.repo.Upsert(ctx, t)
	case MultiFamilyProperty:
		if _, ok := up.(UpdateMultiFamilyProperty); !ok {
			return nil, UpdateWrongPropertyTypeError{Property: saved, UpdateProperty: up}
		}
		if err := mapstructure.Decode(up, &t); err != nil {
			return nil, err
		}
		t.DateUpdated = time.Now().UTC()
		return s.repo.Upsert(ctx, t)
	default:
		err := UnhandledPropertyTypeError{Property: t}
		return nil, fmt.Errorf("unhandled saved property type: %w", err)
	}
}

// Delete removes the Property identified by id from the Repository.
func (s *Service) Delete(ctx context.Context, id uuid.UUID) (Property, error) {
	return s.repo.Delete(ctx, id)
}
