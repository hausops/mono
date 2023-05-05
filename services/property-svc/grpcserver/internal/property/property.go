// Package property implements pb.PropertyServer.
package property

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

type server struct {
	pb.UnimplementedPropertyServer
	svc *property.Service
}

// TODO: handle c.Proxy once adapter/dapr is implemented.
func NewServer(_ config.Config) *server {
	repo := local.
		NewPropertyRepository().
		ReplaceProperties(local.ExampleProperties())

	return &server{svc: property.NewService(repo)}
}

func (s *server) Create(ctx context.Context, in *pb.CreatePropertyRequest) (*pb.PropertyResponse, error) {
	p, err := decodeCreatePropertyRequest(in)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("decode CreatePropertyRequest: %v", err),
		)
	}

	created, err := s.svc.Create(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("property.Service.Create(%v): %w", p, err)
	}
	return newPropertyResponse(created).encode(), nil
}

func (s *server) FindByID(ctx context.Context, in *pb.PropertyIDRequest) (*pb.PropertyResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid property id")
	}

	p, err := s.svc.FindByID(ctx, id)
	if err != nil {
		switch err {
		case property.ErrInvalidID:
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("id=%s", id))
		case property.ErrNotFound:
			return nil, status.Error(codes.NotFound, "Property not found")
		default:
			return nil, fmt.Errorf("find property (id=%s): %w", id, err)
		}
	}
	return newPropertyResponse(p).encode(), nil
}

func (s *server) List(ctx context.Context, _ *emptypb.Empty) (*pb.PropertyListResponse, error) {
	properties, err := s.svc.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list property: %w", err)
	}
	ps := make([]*pb.PropertyResponse, len(properties))
	for i, p := range properties {
		ps[i] = newPropertyResponse(p).encode()
	}
	return &pb.PropertyListResponse{Properties: ps}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdatePropertyRequest) (*pb.PropertyResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid property id")
	}

	up, err := decodeUpdatePropertyRequest(in)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("decode UpdatePropertyRequest: %v", err),
		)
	}

	updated, err := s.svc.Update(ctx, id, up)
	if err != nil {
		switch err {
		case property.ErrInvalidID:
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("id=%s", id))
		case property.ErrNotFound:
			return nil, status.Error(codes.NotFound, "Property not found")
		default:
			return nil, fmt.Errorf("property.Service.Update(%v): %w", in, err)
		}
	}
	return newPropertyResponse(updated).encode(), nil
}

func (s *server) Delete(ctx context.Context, in *pb.PropertyIDRequest) (*pb.PropertyResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid property id")
	}

	p, err := s.svc.Delete(ctx, id)
	if err != nil {
		switch err {
		case property.ErrNotFound:
			return nil, status.Error(codes.NotFound, "Property not found")
		default:
			return nil, fmt.Errorf("delete property (id=%s): %w", id, err)
		}
	}
	return newPropertyResponse(p).encode(), nil
}
