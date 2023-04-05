package property

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (Property, error)
	List(ctx context.Context) ([]Property, error)
}
