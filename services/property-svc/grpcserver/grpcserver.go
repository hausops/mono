package grpcserver

import (
	"context"
	"fmt"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/grpcserver/internal/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

type server struct {
	*grpc.Server
	deps   *dependencies
	logger *zap.Logger
}

func New(ctx context.Context, c config.Config, logger *zap.Logger) (*server, error) {
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.MaxConcurrentStreams(100),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	deps, err := newDependencies(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("new dependencies: %w", err)
	}
	pb.RegisterPropertyServer(s, property.NewServer(deps.propertySvc))

	if c.Mode == config.ModeDev {
		reflection.Register(s)
	}

	return &server{s, deps, logger}, nil
}

// GracefulStop gracefully stops the server and cleans up dependencies.
// If the server doesn't stop within the given timeout via ctx, it forcefully
// stops the server.
func (s *server) GracefulStop(ctx context.Context) {
	stopped := make(chan struct{})
	go func() {
		s.Server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.logger.Info("Server gracefully stopped")
	case <-ctx.Done():
		s.Server.Stop()
		s.logger.Warn("Server forcefully stopped due to timeout exceeded")
	}

	s.logger.Info("Cleaning up dependencies")
	if err := s.deps.close(ctx); err != nil {
		s.logger.Error("Error cleaning up dependencies", zap.Error(err))
	} else {
		s.logger.Info("Successfully cleaned up dependencies")
	}
}
