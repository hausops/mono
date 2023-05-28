package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/adapter/mongo"
	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/domain/property"
	"golang.org/x/sync/errgroup"
)

type dependencies struct {
	propertySvc   *property.Service
	closeHandlers []func(context.Context) error
}

func newDependencies(ctx context.Context, c config.Config) (*dependencies, error) {
	var deps dependencies
	switch t := c.Datastore.(type) {
	case config.LocalDatastore:
		propertyRepo := local.
			NewPropertyRepository().
			ReplaceProperties(local.ExampleProperties())

		deps = dependencies{
			propertySvc: property.NewService(propertyRepo),
		}

	case config.MongoDatastore:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		c, err := mongo.Conn(ctx, t.URI)
		if err != nil {
			return nil, fmt.Errorf("connect to mongo: %w", err)
		}

		deps.onClose(func(ctx context.Context) error {
			return c.Disconnect(ctx)
		})

		propertyRepo := mongo.NewPropertyRepository(c)

		deps = dependencies{
			propertySvc: property.NewService(propertyRepo),
		}
	}

	if err := deps.validate(); err != nil {
		return nil, fmt.Errorf("invalid dependencies: %w", err)
	}

	return &deps, nil
}

func (d *dependencies) validate() error {
	if d.propertySvc == nil {
		return errors.New("property service is not set")
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
