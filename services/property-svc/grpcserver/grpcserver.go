package grpcserver

import (
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"

	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/grpcserver/internal/property"
	"github.com/hausops/mono/services/property-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(c config.Config, logger *zap.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.MaxConcurrentStreams(100),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	deps := newDependencies(c)
	pb.RegisterPropertyServer(s, property.NewServer(deps.propertySvc))

	if c.Mode == config.ModeDev {
		reflection.Register(s)
	}

	return s
}
