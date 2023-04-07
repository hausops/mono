package property

import (
	"fmt"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

// "encode" transforms domain types _to_ pb for transport.
// "decode" transforms _from_ pb to domain types.

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
	addr := address{in.GetAddress()}
	return property.SingleFamilyProperty{
		Address:       addr.decode(),
		CoverImageUrl: in.GetCoverImageUrl(),
		YearBuilt:     in.GetYearBuilt(),
	}
}

type multiFamilyPropertyRequest struct {
	*pb.PropertyRequest_MultiFamilyProperty
}

func (r multiFamilyPropertyRequest) decode() property.MultiFamilyProperty {
	in := r.MultiFamilyProperty
	addr := address{in.GetAddress()}
	return property.MultiFamilyProperty{
		Address:       addr.decode(),
		CoverImageUrl: in.GetCoverImageUrl(),
		YearBuilt:     in.GetYearBuilt(),
	}
}

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
			Id:            p.ID,
			Address:       new(address).encode(p.Address),
			CoverImageUrl: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		},
	}
	return (*pb.PropertyResponse_SingleFamilyProperty)(r)
}

type multiFamilyPropertyResponse pb.PropertyResponse_MultiFamilyProperty

func (r *multiFamilyPropertyResponse) encode(p property.MultiFamilyProperty) *pb.PropertyResponse_MultiFamilyProperty {
	r = &multiFamilyPropertyResponse{
		MultiFamilyProperty: &pb.MultiFamilyProperty{
			Id:            p.ID,
			Address:       new(address).encode(p.Address),
			CoverImageUrl: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		},
	}
	return (*pb.PropertyResponse_MultiFamilyProperty)(r)
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

func (a address) decode() property.Address {
	return property.Address{
		Line1: a.GetLine1(),
		Line2: a.GetLine2(),
		City:  a.GetCity(),
		State: a.GetState(),
		Zip:   a.GetZip(),
	}
}
