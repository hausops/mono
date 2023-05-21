package grpcserver

import (
	"context"

	"github.com/hausops/mono/services/user-svc/config"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type dependencies struct {
	userSvc *user.Service
	cleanUp func(context.Context) error
}

func newDependencies(ctx context.Context, c config.Config) (*dependencies, error) {
	return &dependencies{
		userSvc: nil,
		cleanUp: func(_ context.Context) error { return nil },
	}, nil
}
