package property

import (
	"context"
	"time"

	"github.com/google/uuid"
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
		return nil, ErrNotFound
	}
	return s.repo.FindByID(ctx, pid)
}

// List returns all Properties from the Repository.
func (s *Service) List(ctx context.Context) ([]Property, error) {
	return s.repo.List(ctx)
}

// Delete removes the Property identified by a given ID from the Repository.
func (s *Service) Delete(ctx context.Context, id string) (Property, error) {
	pid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return s.repo.Delete(ctx, pid)
}
