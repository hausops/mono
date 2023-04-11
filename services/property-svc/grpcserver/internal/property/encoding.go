package property

import (
	"fmt"
	"time"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

// "encode" transforms domain types _to_ pb for transport.
// "decode" transforms _from_ pb to domain types.

type propertyResponse pb.PropertyResponse

func (r *propertyResponse) encode(p property.Property) *pb.PropertyResponse {
	switch t := p.(type) {
	case property.SingleFamilyProperty:
		r = &propertyResponse{
			Property: new(singleFamilyPropertyResponse).encode(t),
		}
	case property.MultiFamilyProperty:
		r = &propertyResponse{
			Property: new(multiFamilyPropertyResponse).encode(t),
		}
	}
	return (*pb.PropertyResponse)(r)
}

type singleFamilyPropertyResponse pb.PropertyResponse_SingleFamilyProperty

func (r *singleFamilyPropertyResponse) encode(p property.SingleFamilyProperty) *pb.PropertyResponse_SingleFamilyProperty {
	r = &singleFamilyPropertyResponse{
		SingleFamilyProperty: &pb.SingleFamilyProperty{
			Id:            p.ID.String(),
			Address:       new(address).encode(p.Address),
			CoverImageUrl: p.CoverImageURL,
			YearBuilt:     p.YearBuilt,
			Unit:          new(rentalUnit).encode(p.Unit),
			DateCreated:   p.DateCreated.Format(time.RFC3339),
		},
	}
	return (*pb.PropertyResponse_SingleFamilyProperty)(r)
}

type multiFamilyPropertyResponse pb.PropertyResponse_MultiFamilyProperty

func (r *multiFamilyPropertyResponse) encode(p property.MultiFamilyProperty) *pb.PropertyResponse_MultiFamilyProperty {
	units := make([]*pb.RentalUnit, len(p.Units))
	for i, u := range p.Units {
		units[i] = new(rentalUnit).encode(u)
	}

	r = &multiFamilyPropertyResponse{
		MultiFamilyProperty: &pb.MultiFamilyProperty{
			Id:            p.ID.String(),
			Address:       new(address).encode(p.Address),
			CoverImageUrl: p.CoverImageURL,
			YearBuilt:     p.YearBuilt,
			Units:         units,
			DateCreated:   p.DateCreated.Format(time.RFC3339),
		},
	}
	return (*pb.PropertyResponse_MultiFamilyProperty)(r)
}

type rentalUnit pb.RentalUnit

func (u *rentalUnit) encode(in property.RentalUnit) *pb.RentalUnit {
	u = &rentalUnit{
		Id:          in.ID.String(),
		Number:      in.Number,
		Bedrooms:    in.Bedrooms,
		Bathrooms:   in.Bathrooms,
		Size:        in.Size,
		RentAmount:  in.RentAmount,
		DateCreated: in.DateCreated.Format(time.RFC3339),
	}
	return (*pb.RentalUnit)(u)
}

type propertyRequest struct {
	*pb.PropertyRequest
}

func (r propertyRequest) decode() property.Property {
	switch t := r.GetProperty().(type) {
	case *pb.PropertyRequest_SingleFamilyProperty:
		sr := singleFamilyPropertyRequest{t}
		return sr.decode()
	case *pb.PropertyRequest_MultiFamilyProperty:
		mr := multiFamilyPropertyRequest{t}
		return mr.decode()
	default:
		// This should never happen (programming error) so we panic.
		panic(fmt.Sprintf("encode propertyRequest: unhandled type %T", t))
	}
}

type singleFamilyPropertyRequest struct {
	*pb.PropertyRequest_SingleFamilyProperty
}

func (r singleFamilyPropertyRequest) decode() property.SingleFamilyProperty {
	in := r.SingleFamilyProperty
	addr := addressRequest{in.GetAddress()}
	unit := rentalUnitRequest{in.GetUnit()}
	return property.SingleFamilyProperty{
		Address:       addr.decode(),
		CoverImageURL: in.GetCoverImageUrl(),
		YearBuilt:     in.GetYearBuilt(),
		Unit:          unit.decode(),
	}
}

type multiFamilyPropertyRequest struct {
	*pb.PropertyRequest_MultiFamilyProperty
}

func (r multiFamilyPropertyRequest) decode() property.MultiFamilyProperty {
	in := r.MultiFamilyProperty
	addr := addressRequest{in.GetAddress()}

	units := make([]property.RentalUnit, len(in.GetUnits()))
	for i, r := range in.GetUnits() {
		u := &rentalUnitRequest{r}
		units[i] = u.decode()
	}

	return property.MultiFamilyProperty{
		Address:       addr.decode(),
		CoverImageURL: in.GetCoverImageUrl(),
		YearBuilt:     in.GetYearBuilt(),
		Units:         units,
	}
}

type rentalUnitRequest struct {
	*pb.RentalUnitRequest
}

func (u rentalUnitRequest) decode() property.RentalUnit {
	return property.RentalUnit{
		Number:     u.GetNumber(),
		Bedrooms:   u.GetBedrooms(),
		Bathrooms:  u.GetBathrooms(),
		Size:       u.GetSize(),
		RentAmount: u.GetRentAmount(),
	}
}

type address struct {
	*pb.Address
}

func (a *address) encode(in property.Address) *pb.Address {
	a = &address{&pb.Address{
		Line1: in.Line1,
		Line2: in.Line2,
		City:  in.City,
		State: in.State,
		Zip:   in.Zip,
	}}
	return a.Address
}

type addressRequest struct {
	*pb.AddressRequest
}

func (a addressRequest) decode() property.Address {
	return property.Address{
		Line1: a.GetLine1(),
		Line2: a.GetLine2(),
		City:  a.GetCity(),
		State: a.GetState(),
		Zip:   a.GetZip(),
	}
}
