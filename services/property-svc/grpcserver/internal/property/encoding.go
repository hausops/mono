package property

import (
	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

func encodeSingleFamiltyProperty(p property.SingleFamilyProperty) *pb.PropertyResponse_SingleFamilyProperty {
	return &pb.PropertyResponse_SingleFamilyProperty{
		SingleFamilyProperty: &pb.SingleFamilyProperty{
			Id:            p.ID,
			Address:       encodeAddress(p.Address),
			CoverImageUrl: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		},
	}
}

func encodeMultiFamiltyProperty(p property.MultiFamilyProperty) *pb.PropertyResponse_MultiFamilyProperty {
	return &pb.PropertyResponse_MultiFamilyProperty{
		MultiFamilyProperty: &pb.MultiFamilyProperty{
			Id:            p.ID,
			Address:       encodeAddress(p.Address),
			CoverImageUrl: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		},
	}
}

func encodeAddress(a property.Address) *pb.Address {
	return &pb.Address{
		Line1: a.Line1,
		Line2: a.Line2,
		City:  a.City,
		State: a.State,
		Zip:   a.Zip,
	}
}
