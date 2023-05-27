package grpcserver

import (
	"context"
	"fmt"
	"time"

	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/adapter/mongo"
	"github.com/hausops/mono/services/user-svc/config"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type dependencies struct {
	userRepo user.Repository
	cleanUp  func(context.Context) error
}

func newDependencies(ctx context.Context, c config.Config) (*dependencies, error) {
	var deps dependencies
	switch t := c.Datastore.(type) {
	case config.LocalDatastore:
		deps = dependencies{
			userRepo: local.NewUserRepository(),
			cleanUp:  func(_ context.Context) error { return nil },
		}

	case config.MongoDatastore:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		c, err := mongo.Conn(ctx, t.URI)
		if err != nil {
			return nil, fmt.Errorf("connect to mongo: %w", err)
		}
		uc := c.Database("user-svc").Collection("users")
		userRepo, err := mongo.NewUserRepository(ctx, uc)
		if err != nil {
			return nil, fmt.Errorf("new user repository (mongo): %w", err)
		}

		deps = dependencies{
			userRepo: userRepo,
			cleanUp: func(ctx context.Context) error {
				return c.Disconnect(ctx)
			},
		}
	}

	return &deps, nil
}
