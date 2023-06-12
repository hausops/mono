package grpcserver

import (
	"context"
	"fmt"

	"github.com/hausops/mono/services/auth-svc/adapter/dapr"
	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/config"
	"github.com/hausops/mono/services/auth-svc/domain/auth"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type dependencies struct {
	authSvc       *auth.Service
	closeHandlers []func(context.Context) error
}

func newDependencies(
	ctx context.Context,
	conf config.Config,
	log *zap.Logger,
) (*dependencies, error) {
	var deps dependencies

	conn, err := dapr.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("create gRPC connection via dapr: %w", err)
	}

	deps.onClose(func(ctx context.Context) error {
		return conn.Close()
	})

	userSvc := dapr.NewUserService(conn)

	deps.authSvc = auth.NewService(
		userSvc,
		auth.Repositories{
			Confirm:    local.NewConfirmRepository(),
			Credential: local.NewCredentialRepository(),
			Session:    local.NewSessionRepository(),
		},
		local.NewEmailDispatcher(),
		log,
	)

	return &deps, nil
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
