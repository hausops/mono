package property

import "context"

type Repository interface {
	// Delete removes property with id. Returns the deleted property or error
	// if not found.
	Delete(ctx context.Context, id string) (Property, error)

	// FindByID returns Property with id. If no Property found for the id,
	// an error is returned.
	FindByID(ctx context.Context, id string) (Property, error)

	// List returns all Property stored in Repository.
	List(ctx context.Context) ([]Property, error)

	// Upsert adds p to Repository if does not exist otherwise it replaces
	// the currently stored Property of the same ID (does not merge).
	Upsert(ctx context.Context, p Property) (Property, error)
}
