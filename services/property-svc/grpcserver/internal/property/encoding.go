package property

import (
	"fmt"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

func encodePropertyResponse(p property.Property) *pb.PropertyResponse {
	switch t := p.(type) {
	case property.SingleFamilyProperty:
		return &pb.PropertyResponse{
			Property: encodeSingleFamiltyProperty(t),
		}
	case property.MultiFamilyProperty:
		return &pb.PropertyResponse{
			Property: encodeMultiFamiltyProperty(t),
		}
	// TODO: move this check to domain layer
	default:
		// This should never happen (programming error) so we panic.
		err := &property.UnhandledPropertyTypeError{Property: t}
		panic(fmt.Errorf("encode PropertyResponse: %w", err))
	}
}

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
