package grpcserver

import (
	"time"

	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/grpcserver/internal/property"
	"github.com/hausops/mono/services/property-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New() *grpc.Server {
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.MaxConcurrentStreams(100),
	)
	pb.RegisterPropertyServer(s, property.NewServer(local.NewPropertyRepository()))

	// TODO only enable in dev mode
	reflection.Register(s)

	return s
}
