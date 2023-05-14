package grpcserver

import (
	"context"
	"fmt"
	"time"

	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/adapter/mongo"
	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

type dependencies struct {
	propertySvc *property.Service
	cleanUp     func(context.Context) error
}

// type dependency interface {
// 	cleanUp(context.Context) error
// }

func newDependencies(ctx context.Context, c config.Config) (*dependencies, error) {
	var deps dependencies
	switch t := c.Datastore.(type) {
	case config.LocalDatastore:
		propertyRepo := local.
			NewPropertyRepository().
			ReplaceProperties(local.ExampleProperties())

		deps = dependencies{
			propertySvc: property.NewService(propertyRepo),
			cleanUp:     func(_ context.Context) error { return nil },
		}

	case config.MongoDatastore:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		mongoClient, err := mongo.Conn(ctx, t.URI)
		if err != nil {
			return nil, fmt.Errorf("connect to mongo: %w", err)
		}
		propertyRepo := mongo.NewPropertyRepository(mongoClient)

		deps = dependencies{
			propertySvc: property.NewService(propertyRepo),
			cleanUp: func(ctx context.Context) error {
				return mongoClient.Disconnect(ctx)
			},
		}
	}

	return &deps, nil
}
