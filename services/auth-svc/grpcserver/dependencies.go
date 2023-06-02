package grpcserver

import (
	"context"

	"github.com/hausops/mono/services/auth-svc/config"
	"golang.org/x/sync/errgroup"
)

type dependencies struct {
	closeHandlers []func(context.Context) error
}

func newDependencies(ctx context.Context, conf config.Config) (*dependencies, error) {
	var deps dependencies
	return &deps, nil
}

// func (d *dependencies) validate() error {
// 	return nil
// }

// func (d *dependencies) onClose(h func(context.Context) error) {
// 	d.closeHandlers = append(d.closeHandlers, h)
// }

func (d *dependencies) close(ctx context.Context) error {
	var g errgroup.Group
	for _, h := range d.closeHandlers {
		h := h // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error { return h(ctx) })
	}
	return g.Wait()
}
