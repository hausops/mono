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

	"github.com/hausops/mono/services/auth-svc/config"
	"github.com/hausops/mono/services/auth-svc/grpcserver/internal/auth"
	"github.com/hausops/mono/services/auth-svc/pb"
)

type server struct {
	*grpc.Server
	deps *dependencies
	log  *zap.Logger
}

func New(ctx context.Context, conf config.Config, log *zap.Logger) (*server, error) {
	srv := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.MaxConcurrentStreams(100),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(log),
		)),
	)

	deps, err := newDependencies(ctx, conf, log)
	if err != nil {
		return nil, fmt.Errorf("new dependencies: %w", err)
	}

	pb.RegisterAuthServer(srv, auth.NewServer(deps.authSvc))

	if conf.Mode == config.ModeDev {
		reflection.Register(srv)
	}

	return &server{srv, deps, log}, nil
}

// GracefulStop gracefully stops the server and cleans up dependencies.
// If the server doesn't stop within the given timeout via ctx, it forcefully
// stops the server.
func (s *server) GracefulStop(ctx context.Context) {
	// flushes logs at the end
	defer s.log.Sync()

	stopped := make(chan struct{})
	go func() {
		s.Server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.log.Info("Server gracefully stopped")
	case <-ctx.Done():
		s.Server.Stop()
		s.log.Warn("Server forcefully stopped due to timeout exceeded")
	}

	s.log.Info("Cleaning up dependencies")
	if err := s.deps.close(ctx); err != nil {
		s.log.Error("Error cleaning up dependencies", zap.Error(err))
	} else {
		s.log.Info("Successfully cleaned up dependencies")
	}
}
