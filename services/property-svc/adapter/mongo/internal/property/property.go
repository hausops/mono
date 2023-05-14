// Package mongo implements property domain backed by mongodb.
package property

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/domain/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		collection: client.Database("property-svc").Collection("properties"),
	}
}

// Ensure repository implements the property.Repository interface.
var _ property.Repository = (*Repository)(nil)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (property.Property, error) {
	res := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id})

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, property.ErrNotFound
		}
		return nil, err
	}

	var raw bson.Raw
	if err := res.Decode(&raw); err != nil {
		return nil, err
	}
	return decodePropertyFromBSON(raw)
}

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (property.Property, error) {
	res := r.collection.FindOne(ctx, bson.M{"_id": id})

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, property.ErrNotFound
		}
		return nil, err
	}

	var raw bson.Raw
	if err := res.Decode(&raw); err != nil {
		return nil, err
	}
	return decodePropertyFromBSON(raw)
}

func (r *Repository) List(ctx context.Context) ([]property.Property, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []bson.Raw
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	ps := make([]property.Property, len(results))
	for i, raw := range results {
		p, err := decodePropertyFromBSON(raw)
		if err != nil {
			return nil, fmt.Errorf("decode property from results[%d]: %w", i, err)
		}
		ps[i] = p
	}
	return ps, nil
}

func (r *Repository) Upsert(ctx context.Context, p property.Property) (property.Property, error) {
	if p == nil {
		return nil, fmt.Errorf("invalid parameter: property is nil")
	}

	if p.GetID() == (uuid.UUID{}) {
		return nil, property.MissingIDError{Message: "missing property ID"}
	}

	switch t := p.(type) {
	case property.SingleFamilyProperty:
		if t.Unit.ID == (uuid.UUID{}) {
			return nil, property.MissingIDError{Message: "missing unit ID"}
		}
	case property.MultiFamilyProperty:
		for _, unit := range t.Units {
			if unit.ID == (uuid.UUID{}) {
				return nil, property.MissingIDError{Message: "missing unit ID"}
			}
		}
	}

	up, err := toPropertyBSON(p)
	if err != nil {
		return nil, fmt.Errorf("convert propety to BSON: %w", err)
	}

	id := p.GetID()
	_, err = r.collection.UpdateByID(ctx, id, bson.M{"$set": up},
		options.Update().SetUpsert(true))

	if err != nil {
		return nil, fmt.Errorf("upsert to collection: %w", err)
	}
	return p, nil
}
