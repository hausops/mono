package property

import (
	"context"
)

type Service struct {
	repo Repository
}

// NewService creates property.service with dependencies.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id string) (Property, error) {
	return s.repo.FindByID(ctx, id)
}
