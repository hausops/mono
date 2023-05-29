package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/adapter/mongo"
	"github.com/hausops/mono/services/user-svc/config"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"golang.org/x/sync/errgroup"
)

type dependencies struct {
	userRepo      user.Repository
	closeHandlers []func(context.Context) error
}

func newDependencies(ctx context.Context, conf config.Config) (*dependencies, error) {
	var deps dependencies
	switch t := conf.Datastore.(type) {
	case config.LocalDatastore:
		deps = dependencies{
			userRepo: local.NewUserRepository(),
		}

	case config.MongoDatastore:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		client, err := mongo.Conn(ctx, t.URI)
		if err != nil {
			return nil, fmt.Errorf("connect to mongo: %w", err)
		}

		deps.onClose(func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})

		uc := client.Database("user-svc").Collection("users")
		userRepo, err := mongo.NewUserRepository(ctx, uc)
		if err != nil {
			return nil, fmt.Errorf("new user repository (mongo): %w", err)
		}

		deps = dependencies{
			userRepo: userRepo,
		}
	}

	if err := deps.validate(); err != nil {
		return nil, fmt.Errorf("invalid dependencies: %w", err)
	}

	return &deps, nil
}

func (d *dependencies) validate() error {
	if d.userRepo == nil {
		return errors.New("user repo is not set")
	}
	return nil
}

func (d *dependencies) onClose(h func(context.Context) error) {
	d.closeHandlers = append(d.closeHandlers, h)
}

func (d *dependencies) close(ctx context.Context) error {
	var g errgroup.Group
	for _, h := range d.closeHandlers {
		h := h // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error { return h(ctx) })
	}
	return g.Wait()
}
