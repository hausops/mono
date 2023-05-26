package user

// Service contains the core logic. It exposes a set of public CRUD APIs.
type Service struct {
	repo Repository
}

// NewService creates a user.Service with its dependencies.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
