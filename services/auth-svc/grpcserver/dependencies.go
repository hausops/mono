package grpcserver

import (
	"context"
	"fmt"

	"github.com/hausops/mono/services/auth-svc/adapter/dapr"
	"github.com/hausops/mono/services/auth-svc/adapter/local"
	"github.com/hausops/mono/services/auth-svc/config"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"golang.org/x/sync/errgroup"
)

type dependencies struct {
	userSvc       userpb.UserServiceClient
	credentialSvc *credential.Service
	confirmSvc    *confirm.Service
	closeHandlers []func(context.Context) error
}

func newDependencies(ctx context.Context, conf config.Config) (*dependencies, error) {
	var deps dependencies

	conn, err := dapr.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("create gRPC connection via dapr: %w", err)
	}

	deps.onClose(func(ctx context.Context) error {
		return conn.Close()
	})

	deps.userSvc = dapr.NewUserService(conn)
	deps.credentialSvc = credential.NewService(local.NewCredentialRepository())
	deps.confirmSvc = confirm.NewService(
		local.NewPendingConfirmationRepository(),
		local.NewConfirmedEmailRepository(),
		local.NewEmailDispatcher(),
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
