package property

import (
	"context"

	"github.com/google/uuid"
)

// Repository interface declares the behavior this package needs to perists and
// retrieve data.
type Repository interface {
	// Delete removes the property with the given id and returns
	// the deleted property, or an error if the property was not found.
	Delete(ctx context.Context, id uuid.UUID) (Property, error)

	// FindByID returns the property with the given id, or an error
	// if the property was not found.
	FindByID(ctx context.Context, id uuid.UUID) (Property, error)

	// List returns all properties stored in the repository.
	List(ctx context.Context) ([]Property, error)

	// Upsert adds p to the repository if it does not exist, or replaces
	// the stored property with the same ID (without merging).
	Upsert(ctx context.Context, p Property) (Property, error)
}
