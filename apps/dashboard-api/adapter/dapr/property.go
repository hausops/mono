package dapr

import (
	"context"
	"log"

	"github.com/hausops/mono/apps/dashboard-api/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PropertyService struct {
	client pb.PropertyClient
}

var _ property.Service = (*PropertyService)(nil)

func NewPropertyService(grpcConn *grpc.ClientConn) *PropertyService {
	return &PropertyService{
		client: pb.NewPropertyClient(grpcConn),
	}
}

func (r *PropertyService) CreateSingleFamilyProperty(in property.CreateSingleFamilyPropertyInput) (*property.SingleFamilyProperty, error) {
	panic("Not implemented - dapr.PropertyService.CreateSingleFamilyProperty")
}

func (r *PropertyService) CreateMultiFamilyProperty(in property.CreateMultiFamilyPropertyInput) (*property.MultiFamilyProperty, error) {
	panic("Not implemented - dapr.PropertyService.CreateMultiFamilyProperty")
}

func (r *PropertyService) FindByID(id string) (property.Property, error) {
	panic("Not implemented - dapr.PropertyService.FindByID")
}

func (r *PropertyService) FindAll() ([]property.Property, error) {
	ctx := context.TODO()
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.List(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("could not list properties:", err)
		return nil, err
	}

	ps := make([]property.Property, len(res.GetProperties()))
	for i, p := range res.GetProperties() {
		switch t := p.GetProperty().(type) {
		case *pb.PropertyResponse_SingleFamilyProperty:
			ps[i] = decodeSingleFamilyProperty(t.SingleFamilyProperty)
		case *pb.PropertyResponse_MultiFamilyProperty:
			ps[i] = decodeMultiFamilyProperty(t.MultiFamilyProperty)
		}
	}
	return ps, nil
}

func decodeSingleFamilyProperty(in *pb.SingleFamilyProperty) property.SingleFamilyProperty {
	p := property.SingleFamilyProperty{
		ID:      in.Id,
		Address: decodeAddress(in.GetAddress()),
		Unit:    decodeSingleFamilyPropertyUnit(in.GetUnit()),
	}

	if v := in.GetCoverImageUrl(); v != "" {
		p.CoverImageURL = &v
	}

	if v := int(in.GetYearBuilt()); v != 0 {
		p.YearBuilt = &v
	}

	return p
}

func decodeSingleFamilyPropertyUnit(in *pb.RentalUnit) property.SingleFamilyPropertyUnit {
	u := property.SingleFamilyPropertyUnit{
		ID: in.GetId(),
	}

	if v := float64(in.GetBedrooms()); v != 0 {
		u.Bedrooms = &v
	}

	if v := float64(in.GetBathrooms()); v != 0 {
		u.Bathrooms = &v
	}

	if v := float64(in.GetSize()); v != 0 {
		u.Size = &v
	}

	if v := float64(in.GetRentAmount()); v != 0 {
		u.RentAmount = &v
	}

	return u
}

func decodeMultiFamilyProperty(in *pb.MultiFamilyProperty) property.MultiFamilyProperty {
	units := make([]property.MultiFamilyPropertyUnit, len(in.GetUnits()))
	for i, u := range in.GetUnits() {
		units[i] = decodeMultiFamilyPropertyUnit(u)
	}

	p := property.MultiFamilyProperty{
		ID:      in.Id,
		Address: decodeAddress(in.GetAddress()),
		Units:   units,
	}

	if v := in.GetCoverImageUrl(); v != "" {
		p.CoverImageURL = &v
	}

	if v := int(in.GetYearBuilt()); v != 0 {
		p.YearBuilt = &v
	}

	return p
}

func decodeMultiFamilyPropertyUnit(in *pb.RentalUnit) property.MultiFamilyPropertyUnit {
	u := property.MultiFamilyPropertyUnit{
		ID:     in.GetId(),
		Number: in.GetNumber(),
	}

	if v := float64(in.GetBedrooms()); v != 0 {
		u.Bedrooms = &v
	}

	if v := float64(in.GetBathrooms()); v != 0 {
		u.Bathrooms = &v
	}

	if v := float64(in.GetSize()); v != 0 {
		u.Size = &v
	}

	if v := float64(in.GetRentAmount()); v != 0 {
		u.RentAmount = &v
	}

	return u
}

func decodeAddress(in *pb.Address) property.Address {
	a := property.Address{
		Line1: in.GetLine1(),
		City:  in.GetCity(),
		State: in.GetState(),
		Zip:   in.GetZip(),
	}

	if v := in.GetLine2(); v != "" {
		a.Line2 = &v
	}

	return a
}

func (r *PropertyService) UpdateSingleFamilyPropertyByID(
	id string,
	in property.UpdateSingleFamilyPropertyInput,
) (*property.SingleFamilyProperty, error) {
	panic("Not implemented - dapr.PropertyService.UpdateSingleFamilyPropertyByID")
}

func (r *PropertyService) UpdateMultiFamilyPropertyByID(
	id string,
	in property.UpdateMultiFamilyPropertyInput,
) (*property.MultiFamilyProperty, error) {
	panic("Not implemented - dapr.PropertyService.UpdateMultiFamilyPropertyByID")
}

func (r *PropertyService) DeleteByID(id string) (property.Property, error) {
	panic("Not implemented - dapr.PropertyService.DeleteByID")
}
