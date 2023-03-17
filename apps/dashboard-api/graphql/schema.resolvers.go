package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/hausops/mono/apps/dashboard-api/domain/property"
)

// CreateSingleFamilyProperty is the resolver for the createSingleFamilyProperty field.
func (r *mutationResolver) CreateSingleFamilyProperty(ctx context.Context, input property.CreateSingleFamilyPropertyInput) (*property.SingleFamilyProperty, error) {
	return r.Property.CreateSingleFamilyProperty(input)
}

// CreateMultiFamilyProperty is the resolver for the createMultiFamilyProperty field.
func (r *mutationResolver) CreateMultiFamilyProperty(ctx context.Context, input property.CreateMultiFamilyPropertyInput) (*property.MultiFamilyProperty, error) {
	return r.Property.CreateMultiFamilyProperty(input)
}

// DeleteProperty is the resolver for the deleteProperty field.
func (r *mutationResolver) DeleteProperty(ctx context.Context, id string) (property.Property, error) {
	return r.Property.DeleteByID(id)
}

// Properties is the resolver for the properties field.
func (r *queryResolver) Properties(ctx context.Context) ([]property.Property, error) {
	return r.Property.FindAll()
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
