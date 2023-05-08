// Package mongo implements domain components backed by mongodb.
package mongo

import (
	"context"
	"time"

	"github.com/hausops/mono/services/property-svc/adapter/mongo/internal/property"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Conn(ctx context.Context) (*mongo.Client, error) {
	uri := "mongodb://localhost:27017"
	opt := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(16).
		SetTimeout(1 * time.Second)

	return mongo.Connect(ctx, opt)
}

func NewPropertyRepository(client *mongo.Client) *property.Repository {
	return property.NewRepository(client)
}
