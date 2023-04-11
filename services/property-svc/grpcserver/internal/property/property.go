// Package property implements pb.PropertyServer.
package property

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

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
	r := propertyRequest{in}
	p := r.decode()
	created, err := s.svc.Create(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("create property (input=%v): %w", p, err)
	}
	return new(propertyResponse).encode(created), nil
}

func (s *server) FindByID(ctx context.Context, in *pb.PropertyIDRequest) (*pb.PropertyResponse, error) {
	id := in.GetId()
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
	return new(propertyResponse).encode(p), nil
}

func (s *server) List(ctx context.Context, _ *emptypb.Empty) (*pb.PropertyListResponse, error) {
	properties, err := s.svc.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list property: %w", err)
	}

	ps := make([]*pb.PropertyResponse, len(properties))
	for i, p := range properties {
		ps[i] = new(propertyResponse).encode(p)
	}
	return &pb.PropertyListResponse{Properties: ps}, nil
}

func (s *server) Update(ctx context.Context, in *pb.PropertyRequest) (*pb.PropertyResponse, error) {
	id := in.GetId()

	var up property.UpdateProperty
	switch t := in.GetProperty().(type) {
	case *pb.PropertyRequest_SingleFamilyProperty:
		p := t.SingleFamilyProperty
		up = property.UpdateProperty{
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		}
		if a := p.Address; a != nil {
			up.Address = &property.UpdateAddress{
				Line1: a.Line1,
				Line2: a.Line2,
				City:  a.City,
				State: a.State,
				Zip:   a.Zip,
			}
		}
	case *pb.PropertyRequest_MultiFamilyProperty:
		p := t.MultiFamilyProperty
		up = property.UpdateProperty{
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		}
		if a := p.Address; a != nil {
			up.Address = &property.UpdateAddress{
				Line1: a.Line1,
				Line2: a.Line2,
				City:  a.City,
				State: a.State,
				Zip:   a.Zip,
			}
		}
	}

	updated, err := s.svc.Update(ctx, id, up)
	if err != nil {
		switch err {
		case property.ErrInvalidID:
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("id=%s", id))
		case property.ErrNotFound:
			return nil, status.Error(codes.NotFound, "Property not found")
		default:
			return nil, fmt.Errorf("update property (input=%v): %w", in, err)
		}
	}
	return new(propertyResponse).encode(updated), nil
}

func (s *server) Delete(ctx context.Context, in *pb.PropertyIDRequest) (*pb.PropertyResponse, error) {
	id := in.GetId()
	p, err := s.svc.Delete(ctx, id)
	if err != nil {
		switch err {
		case property.ErrNotFound:
			return nil, status.Error(codes.NotFound, "Property not found")
		default:
			return nil, fmt.Errorf("delete property (id=%s): %w", id, err)
		}
	}
	return new(propertyResponse).encode(p), nil
}
