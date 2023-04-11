package property

import (
	"fmt"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

// Decoders transform protobuf messages (generally, requests) to domain types.

type createPropertyRequest struct {
	*pb.CreatePropertyRequest
}

func (r createPropertyRequest) decode() property.Property {
	decodeAddress := func(r *pb.CreatePropertyRequest_Address) property.Address {
		return property.Address{
			Line1: r.GetLine1(),
			Line2: r.GetLine2(),
			City:  r.GetCity(),
			State: r.GetState(),
			Zip:   r.GetZip(),
		}
	}

	decodeUnit := func(r *pb.CreatePropertyRequest_RentalUnit) property.RentalUnit {
		return property.RentalUnit{
			Number:     r.GetNumber(),
			Bedrooms:   r.GetBedrooms(),
			Bathrooms:  r.GetBathrooms(),
			Size:       r.GetSize(),
			RentAmount: r.GetRentAmount(),
		}
	}

	if p := r.GetSingleFamilyProperty(); p != nil {
		return property.SingleFamilyProperty{
			Address:       decodeAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
			Unit:          decodeUnit(p.Unit),
		}
	}

	if p := r.GetMultiFamilyProperty(); p != nil {
		units := make([]property.RentalUnit, len(p.GetUnits()))
		for i, r := range p.GetUnits() {
			units[i] = decodeUnit(r)
		}

		return property.MultiFamilyProperty{
			Address:       decodeAddress(p.Address),
			CoverImageURL: p.GetCoverImageUrl(),
			YearBuilt:     p.GetYearBuilt(),
			Units:         units,
		}
	}

	// This should never happen (programming error) so we panic.
	panic(fmt.Sprintf("decode CreatePropertyRequest: unhandled type[%T]", r.GetProperty()))
}

type updatePropertyRequest struct {
	*pb.UpdatePropertyRequest
}

func (r updatePropertyRequest) decode() property.UpdateProperty {
	decodeAddress := func(r *pb.UpdatePropertyRequest_Address) *property.UpdateAddress {
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

	decodeUnit := func(r *pb.UpdatePropertyRequest_RentalUnit) *property.UpdateRentalUnit {
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

	if p := r.GetSingleFamilyProperty(); p != nil {
		return property.UpdateSingleFamilyProperty{
			Address:       decodeAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
			Unit:          decodeUnit(p.Unit),
		}
	}

	if p := r.GetMultiFamilyProperty(); p != nil {
		return property.UpdateMultiFamilyProperty{
			Address:       decodeAddress(p.Address),
			CoverImageURL: p.CoverImageUrl,
			YearBuilt:     p.YearBuilt,
		}
	}

	// This should never happen (programming error) so we panic.
	panic(fmt.Sprintf("decode UpdatePropertyRequest: unhandled type[%T]", r.GetProperty()))
}
