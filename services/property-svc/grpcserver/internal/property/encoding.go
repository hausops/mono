package property

import (
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

// func decodeMultiFamilyProperty(in *pb.PropertyRequest_MultiFamilyProperty) property.MultiFamilyProperty {
// 	if in == nil {
// 		return property.MultiFamilyProperty{}
// 	}
// 	p := in.MultiFamilyProperty
// 	return property.MultiFamilyProperty{
// 		Address:       decodeAddress(p.GetAddress()),
// 		CoverImageUrl: p.GetCoverImageUrl(),
// 		YearBuilt:     p.GetYearBuilt(),
// 	}
// }

type address pb.Address

func (a *address) encode(in property.Address) *pb.Address {
	a = &address{
		Line1: in.Line1,
		Line2: in.Line2,
		City:  in.City,
		State: in.State,
		Zip:   in.Zip,
	}
	return (*pb.Address)(a)
}

// func (a address) decode() property.Address {
// 	return property.Address{
// 		Line1: a.GetLine1(),
// 		Line2: a.GetLine2(),
// 		City:  a.GetCity(),
// 		State: a.GetState(),
// 		Zip:   a.GetZip(),
// 	}
// }
