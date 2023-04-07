package property

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

// NewService creates property.service with dependencies.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, p Property) (Property, error) {
	switch t := p.(type) {
	case SingleFamilyProperty:
		t.ID = uuid.New().String()
		return s.repo.Upsert(ctx, t)
	case MultiFamilyProperty:
		t.ID = uuid.New().String()
		return s.repo.Upsert(ctx, t)
	default:
		return nil, &UnhandledPropertyTypeError{Property: t}
	}
}

func (s *Service) FindByID(ctx context.Context, id string) (Property, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]Property, error) {
	return s.repo.List(ctx)
}
