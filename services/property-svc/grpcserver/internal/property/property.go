// Package property implements pb.PropertyServer.
package property

import (
	"context"
	"fmt"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

type server struct {
	pb.UnimplementedPropertyServer
	svc *property.Service
}

func NewServer(repo property.Repository) *server {
	return &server{svc: property.NewService(repo)}
}

func (s *server) Create(ctx context.Context, in *pb.PropertyRequest) (*pb.PropertyResponse, error) {
	r := propertyRequest{protoMessage: in}
	p := r.decode()
	created, err := s.svc.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return new(propertyResponse).encode(created), nil
}

func (s *server) FindByID(ctx context.Context, in *pb.PropertyIDRequest) (*pb.PropertyResponse, error) {
	id := in.GetId()
	p, err := s.svc.FindByID(ctx, id)
	if err != nil {
		switch err {
		case property.ErrNotFound:
			return nil, fmt.Errorf("no property with id %s", id)
		default:
			return nil, fmt.Errorf("cannot find property: %v", err)
		}
	}

	return new(propertyResponse).encode(p), nil
}

func (s *server) List(ctx context.Context, _ *pb.Empty) (*pb.PropertyListResponse, error) {
	properties, err := s.svc.List(ctx)
	if err != nil {
		return nil, err
	}

	ps := make([]*pb.PropertyResponse, len(properties))
	for i, p := range properties {
		ps[i] = new(propertyResponse).encode(p)
	}
	return &pb.PropertyListResponse{Properties: ps}, nil
}
