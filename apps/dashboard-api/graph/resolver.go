package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/hausops/mono/apps/dashboard-api/domain/property"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	propertyRepo property.Repository
}

func NewResolver(propertyRepo property.Repository) *Resolver {
	return &Resolver{
		propertyRepo: propertyRepo,
	}
}
