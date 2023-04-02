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

	switch p := p.(type) {
	case property.SingleFamilyProperty:
		return &pb.PropertyResponse{
			Property: encodeSingleFamiltyProperty(p),
		}, nil
	case property.MultiFamilyProperty:
		return &pb.PropertyResponse{
			Property: encodeMultiFamiltyProperty(p),
		}, nil
	default:
		panic(fmt.Sprintf("find property by id %s: unhandled type %T", id, p))
	}
}

func (s *server) List(_ *pb.Empty, _ pb.Property_ListServer) error {
	panic("not implemented") // TODO: Implement
}

func (s *server) CreateSingleFamilyProperty(ctx context.Context, in *pb.SingleFamilyPropertyRequest) (*pb.SingleFamilyProperty, error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) CreateMultiFamilyProperty(ctx context.Context, in *pb.MultiFamilyPropertyRequest) (*pb.MultiFamilyProperty, error) {
	panic("not implemented") // TODO: Implement
}

// var _ pb.PropertyServer = (*propertyService)(nil)
