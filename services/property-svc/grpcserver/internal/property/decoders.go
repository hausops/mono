package property

import (
	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

// Decoders transform protobuf messages (generally, requests) to domain types.

func decodeCreatePropertyRequest(r *pb.CreatePropertyRequest) (property.Property, error) {
	switch t := r.GetProperty().(type) {

	case *pb.CreatePropertyRequest_SingleFamilyProperty_:
		p := t.SingleFamilyProperty
		return property.SingleFamilyProperty{
			Address:       decodeCreateAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
			Unit:          decodeCreateUnit(p.Unit),
		}, nil

	case *pb.CreatePropertyRequest_MultiFamilyProperty_:
		p := t.MultiFamilyProperty
		units := make([]property.RentalUnit, len(p.GetUnits()))
		for i, r := range p.GetUnits() {
			units[i] = decodeCreateUnit(r)
		}

		return property.MultiFamilyProperty{
			Address:       decodeCreateAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
			Units:         units,
		}, nil

	default:
		return nil, unsupportedRequestTypeError{Request: t}
	}
}

func decodeCreateAddress(r *pb.CreatePropertyRequest_Address) property.Address {
	return property.Address{
		Line1: r.GetLine1(),
		Line2: r.GetLine2(),
		City:  r.GetCity(),
		State: r.GetState(),
		Zip:   r.GetZip(),
	}
}

func decodeCreateUnit(r *pb.CreatePropertyRequest_RentalUnit) property.RentalUnit {
	return property.RentalUnit{
		Number:     r.GetNumber(),
		Bedrooms:   r.GetBedrooms(),
		Bathrooms:  r.GetBathrooms(),
		Size:       r.GetSize(),
		RentAmount: r.GetRentAmount(),
	}
}

func decodeUpdatePropertyRequest(r *pb.UpdatePropertyRequest) (property.UpdateProperty, error) {
	switch t := r.GetProperty().(type) {

	case *pb.UpdatePropertyRequest_SingleFamilyProperty:
		p := t.SingleFamilyProperty
		return property.UpdateSingleFamilyProperty{
			Address:       decodeUpdateAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
			Unit:          decodeUpdateUnit(p.Unit),
		}, nil

	case *pb.UpdatePropertyRequest_MultiFamilyProperty:
		p := t.MultiFamilyProperty
		return property.UpdateMultiFamilyProperty{
			Address:       decodeUpdateAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		}, nil

	default:
		return nil, unsupportedRequestTypeError{Request: t}
	}
}

func decodeUpdateAddress(r *pb.UpdatePropertyRequest_Address) *property.UpdateAddress {
	if r == nil {
		return nil
	}
	return &property.UpdateAddress{
		Line1: r.Line1,
		Line2: r.Line2,
		City:  r.City,
		State: r.State,
		Zip:   r.Zip,
	}
}

func decodeUpdateUnit(r *pb.UpdatePropertyRequest_RentalUnit) *property.UpdateRentalUnit {
	if r == nil {
		return nil
	}
	return &property.UpdateRentalUnit{
		Number:     r.Number,
		Bedrooms:   r.Bedrooms,
		Bathrooms:  r.Bathrooms,
		Size:       r.Size,
		RentAmount: r.RentAmount,
	}
}
