package dapr

import (
	"context"
	"errors"
	"fmt"

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

func (r *PropertyService) CreateSingleFamilyProperty(
	ctx context.Context,
	in property.CreateSingleFamilyPropertyInput,
) (*property.SingleFamilyProperty, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.Create(ctx, encodeCreateSingleFamilyPropertyInput(in))
	if err != nil {
		return nil, fmt.Errorf("property-svc.Create: %w", err)
	}

	p := res.GetSingleFamilyProperty()
	if p == nil {
		return nil, errors.New("response is not SingleFamilyProperty")
	}
	return decodeSingleFamilyProperty(p), nil
}

func (r *PropertyService) CreateMultiFamilyProperty(
	ctx context.Context,
	in property.CreateMultiFamilyPropertyInput,
) (*property.MultiFamilyProperty, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.Create(ctx, encodeCreateMultiFamilyPropertyInput(in))
	if err != nil {
		return nil, fmt.Errorf("property-svc.Create: %w", err)
	}

	p := res.GetMultiFamilyProperty()
	if p == nil {
		return nil, errors.New("response is not MultiFamilyProperty")
	}
	return decodeMultiFamilyProperty(p), nil
}

func (r *PropertyService) FindByID(ctx context.Context, id string) (property.Property, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.FindByID(ctx, &pb.PropertyIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("propert-svc.FindByID(id: %s): %w", id, err)
	}
	return decodeProperty(res)
}

func (r *PropertyService) FindAll(ctx context.Context) ([]property.Property, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.List(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("propert-svc.List: %w", err)
	}

	ps := make([]property.Property, len(res.GetProperties()))
	for i, p := range res.GetProperties() {
		d, err := decodeProperty(p)
		if err != nil {
			return nil, fmt.Errorf("decode property at index %d: %w", i, err)
		}
		ps[i] = d
	}
	return ps, nil
}

func (r *PropertyService) UpdateSingleFamilyPropertyByID(
	ctx context.Context,
	id string,
	in property.UpdateSingleFamilyPropertyInput,
) (*property.SingleFamilyProperty, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.Update(ctx, encodeUpdateSingleFamilyPropertyInput(id, in))
	if err != nil {
		return nil, fmt.Errorf("property-svc.Update(id: %s): %w", id, err)
	}

	p := res.GetSingleFamilyProperty()
	if p == nil {
		return nil, errors.New("response is not SingleFamilyProperty")
	}
	return decodeSingleFamilyProperty(p), nil
}

func (r *PropertyService) UpdateMultiFamilyPropertyByID(
	ctx context.Context,
	id string,
	in property.UpdateMultiFamilyPropertyInput,
) (*property.MultiFamilyProperty, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.Update(ctx, encodeUpdateMultiFamilyPropertyInput(id, in))
	if err != nil {
		return nil, fmt.Errorf("property-svc.Update(id: %s): %w", id, err)
	}

	p := res.GetMultiFamilyProperty()
	if p == nil {
		return nil, errors.New("response is not MultiFamilyProperty")
	}
	return decodeMultiFamilyProperty(p), nil
}

func (r *PropertyService) DeleteByID(ctx context.Context, id string) (property.Property, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "property-svc")

	res, err := r.client.Delete(ctx, &pb.PropertyIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("property-svc.Delete(id: %s): %w", id, err)
	}
	return decodeProperty(res)
}

func decodeProperty(r *pb.PropertyResponse) (property.Property, error) {
	switch t := r.GetProperty().(type) {
	case *pb.PropertyResponse_SingleFamilyProperty:
		return decodeSingleFamilyProperty(t.SingleFamilyProperty), nil
	case *pb.PropertyResponse_MultiFamilyProperty:
		return decodeMultiFamilyProperty(t.MultiFamilyProperty), nil
	default:
		return nil, fmt.Errorf("unsupported property type: %T", t)
	}
}

func decodeSingleFamilyProperty(r *pb.SingleFamilyProperty) *property.SingleFamilyProperty {
	p := property.SingleFamilyProperty{
		ID:      r.Id,
		Address: decodeAddress(r.GetAddress()),
		Unit:    decodeSingleFamilyPropertyUnit(r.GetUnit()),
	}

	if v := r.GetCoverImageUrl(); v != "" {
		p.CoverImageURL = &v
	}

	if v := int(r.GetYearBuilt()); v != 0 {
		p.YearBuilt = &v
	}

	return &p
}

func decodeMultiFamilyProperty(r *pb.MultiFamilyProperty) *property.MultiFamilyProperty {
	units := make([]property.MultiFamilyPropertyUnit, len(r.GetUnits()))
	for i, u := range r.GetUnits() {
		units[i] = decodeMultiFamilyPropertyUnit(u)
	}

	p := property.MultiFamilyProperty{
		ID:      r.Id,
		Address: decodeAddress(r.GetAddress()),
		Units:   units,
	}

	if v := r.GetCoverImageUrl(); v != "" {
		p.CoverImageURL = &v
	}

	if v := int(r.GetYearBuilt()); v != 0 {
		p.YearBuilt = &v
	}

	return &p
}

func decodeAddress(r *pb.Address) property.Address {
	a := property.Address{
		Line1: r.GetLine1(),
		City:  r.GetCity(),
		State: r.GetState(),
		Zip:   r.GetZip(),
	}

	if v := r.GetLine2(); v != "" {
		a.Line2 = &v
	}

	return a
}

func decodeSingleFamilyPropertyUnit(r *pb.RentalUnit) property.SingleFamilyPropertyUnit {
	u := property.SingleFamilyPropertyUnit{
		ID: r.GetId(),
	}

	if v := float64(r.GetBedrooms()); v != 0 {
		u.Bedrooms = &v
	}

	if v := float64(r.GetBathrooms()); v != 0 {
		u.Bathrooms = &v
	}

	if v := float64(r.GetSize()); v != 0 {
		u.Size = &v
	}

	if v := float64(r.GetRentAmount()); v != 0 {
		u.RentAmount = &v
	}

	return u
}

func decodeMultiFamilyPropertyUnit(r *pb.RentalUnit) property.MultiFamilyPropertyUnit {
	u := property.MultiFamilyPropertyUnit{
		ID:     r.GetId(),
		Number: r.GetNumber(),
	}

	if v := float64(r.GetBedrooms()); v != 0 {
		u.Bedrooms = &v
	}

	if v := float64(r.GetBathrooms()); v != 0 {
		u.Bathrooms = &v
	}

	if v := float64(r.GetSize()); v != 0 {
		u.Size = &v
	}

	if v := float64(r.GetRentAmount()); v != 0 {
		u.RentAmount = &v
	}

	return u
}

func encodeCreateSingleFamilyPropertyInput(
	in property.CreateSingleFamilyPropertyInput,
) *pb.CreatePropertyRequest {
	p := &pb.CreatePropertyRequest_SingleFamilyProperty{
		Address: encodeCreateAddressInput(in.Address),
		Unit:    encodeCreateSingleFamilyPropertyUnitInput(in.Unit),
	}

	if in.CoverImageURL != nil {
		p.CoverImageUrl = *in.CoverImageURL
	}

	if in.YearBuilt != nil {
		p.YearBuilt = int32(*in.YearBuilt)
	}

	return &pb.CreatePropertyRequest{
		Property: &pb.CreatePropertyRequest_SingleFamilyProperty_{
			SingleFamilyProperty: p,
		},
	}
}

func encodeCreateMultiFamilyPropertyInput(
	in property.CreateMultiFamilyPropertyInput,
) *pb.CreatePropertyRequest {
	units := make([]*pb.CreatePropertyRequest_RentalUnit, len(in.Units))
	for i, u := range in.Units {
		units[i] = encodeCreateMultiFamilyPropertyUnitInput(u)
	}

	p := &pb.CreatePropertyRequest_MultiFamilyProperty{
		Address: encodeCreateAddressInput(in.Address),
		Units:   units,
	}

	if in.CoverImageURL != nil {
		p.CoverImageUrl = *in.CoverImageURL
	}

	if in.YearBuilt != nil {
		p.YearBuilt = int32(*in.YearBuilt)
	}

	return &pb.CreatePropertyRequest{
		Property: &pb.CreatePropertyRequest_MultiFamilyProperty_{
			MultiFamilyProperty: p,
		},
	}
}

func encodeCreateAddressInput(
	in property.CreateAddressInput,
) *pb.CreatePropertyRequest_Address {
	a := &pb.CreatePropertyRequest_Address{
		Line1: in.Line1,
		City:  in.City,
		State: in.State,
		Zip:   in.Zip,
	}

	if in.Line2 != nil {
		a.Line2 = *in.Line2
	}

	return a
}

func encodeCreateSingleFamilyPropertyUnitInput(
	in property.CreateSingleFamilyPropertyUnitInput,
) *pb.CreatePropertyRequest_RentalUnit {
	u := new(pb.CreatePropertyRequest_RentalUnit)

	if in.Bedrooms != nil {
		u.Bedrooms = float32(*in.Bedrooms)
	}

	if in.Bathrooms != nil {
		u.Bathrooms = float32(*in.Bathrooms)
	}

	if in.Size != nil {
		u.Size = float32(*in.Size)
	}

	if in.RentAmount != nil {
		u.RentAmount = float32(*in.RentAmount)
	}

	return u
}

func encodeCreateMultiFamilyPropertyUnitInput(
	in property.CreateMultiFamilyPropertyUnitInput,
) *pb.CreatePropertyRequest_RentalUnit {
	u := &pb.CreatePropertyRequest_RentalUnit{
		Number: in.Number,
	}

	if in.Bedrooms != nil {
		u.Bedrooms = float32(*in.Bedrooms)
	}

	if in.Bathrooms != nil {
		u.Bathrooms = float32(*in.Bathrooms)
	}

	if in.Size != nil {
		u.Size = float32(*in.Size)
	}

	if in.RentAmount != nil {
		u.RentAmount = float32(*in.RentAmount)
	}

	return u
}

func encodeUpdateSingleFamilyPropertyInput(
	id string,
	in property.UpdateSingleFamilyPropertyInput,
) *pb.UpdatePropertyRequest {
	p := &pb.UpdatePropertyRequest_UpdateSingleFamilyProperty{
		Address:       encodeUpdateAddressInput(in.Address),
		CoverImageUrl: in.CoverImageURL,
		Unit:          encodeUpdateSingleFamilyPropertyUnitInput(in.Unit),
	}

	if in.YearBuilt != nil {
		v := int32(*in.YearBuilt)
		p.YearBuilt = &v
	}

	return &pb.UpdatePropertyRequest{
		Property: &pb.UpdatePropertyRequest_SingleFamilyProperty{
			SingleFamilyProperty: p,
		},
	}
}

func encodeUpdateMultiFamilyPropertyInput(
	id string,
	in property.UpdateMultiFamilyPropertyInput,
) *pb.UpdatePropertyRequest {
	p := &pb.UpdatePropertyRequest_UpdateMultiFamilyProperty{
		Address:       encodeUpdateAddressInput(in.Address),
		CoverImageUrl: in.CoverImageURL,
	}

	if in.YearBuilt != nil {
		v := int32(*in.YearBuilt)
		p.YearBuilt = &v
	}

	return &pb.UpdatePropertyRequest{
		Property: &pb.UpdatePropertyRequest_MultiFamilyProperty{
			MultiFamilyProperty: p,
		},
	}
}

func encodeUpdateAddressInput(
	in *property.UpdateAddressInput,
) *pb.UpdatePropertyRequest_Address {
	if in == nil {
		return nil
	}
	return &pb.UpdatePropertyRequest_Address{
		Line1: in.Line1,
		Line2: in.Line2,
		City:  in.City,
		State: in.State,
		Zip:   in.Zip,
	}
}

func encodeUpdateSingleFamilyPropertyUnitInput(
	in *property.UpdateSingleFamilyPropertyUnitInput,
) *pb.UpdatePropertyRequest_RentalUnit {
	if in == nil {
		return nil
	}

	u := new(pb.UpdatePropertyRequest_RentalUnit)

	if in.Bedrooms != nil {
		v := float32(*in.Bedrooms)
		u.Bedrooms = &v
	}

	if in.Bathrooms != nil {
		v := float32(*in.Bathrooms)
		u.Bathrooms = &v
	}

	if in.Size != nil {
		v := float32(*in.Size)
		u.Size = &v
	}

	if in.RentAmount != nil {
		v := float32(*in.RentAmount)
		u.RentAmount = &v
	}

	return u
}
