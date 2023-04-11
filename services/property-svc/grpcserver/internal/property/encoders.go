package property

import (
	"time"

	"github.com/hausops/mono/services/property-svc/domain/property"
	"github.com/hausops/mono/services/property-svc/pb"
)

// Encoders transform domain types to protobuf messages (generally, responses).

type propertyResponse struct {
	property.Property
}

func newPropertyResponse(p property.Property) propertyResponse {
	return propertyResponse{Property: p}
}

func (r propertyResponse) encode() *pb.PropertyResponse {
	switch t := r.Property.(type) {
	case property.SingleFamilyProperty:
		return &pb.PropertyResponse{
			Property: singleFamilyPropertyResponse{t}.encode(),
		}
	case property.MultiFamilyProperty:
		return &pb.PropertyResponse{
			Property: multiFamilyPropertyResponse{t}.encode(),
		}
	default:
		// This should never happen (programming error) so we panic.
		panic(property.UnhandledPropertyTypeError{Property: t})
	}
}

type singleFamilyPropertyResponse struct {
	property.SingleFamilyProperty
}

func (r singleFamilyPropertyResponse) encode() *pb.PropertyResponse_SingleFamilyProperty {
	return &pb.PropertyResponse_SingleFamilyProperty{
		SingleFamilyProperty: &pb.SingleFamilyProperty{
			Id:            r.ID.String(),
			Address:       address{r.Address}.encode(),
			CoverImageUrl: r.CoverImageURL,
			YearBuilt:     r.YearBuilt,
			Unit:          rentalUnit{r.Unit}.encode(),
			DateCreated:   r.DateCreated.Format(time.RFC3339),
			DateUpdated:   r.DateUpdated.Format(time.RFC3339),
		},
	}
}

type multiFamilyPropertyResponse struct {
	property.MultiFamilyProperty
}

func (r multiFamilyPropertyResponse) encode() *pb.PropertyResponse_MultiFamilyProperty {
	units := make([]*pb.RentalUnit, len(r.Units))
	for i, u := range r.Units {
		units[i] = rentalUnit{u}.encode()
	}

	return &pb.PropertyResponse_MultiFamilyProperty{
		MultiFamilyProperty: &pb.MultiFamilyProperty{
			Id:            r.ID.String(),
			Address:       address{r.Address}.encode(),
			CoverImageUrl: r.CoverImageURL,
			YearBuilt:     r.YearBuilt,
			Units:         units,
			DateCreated:   r.DateCreated.Format(time.RFC3339),
			DateUpdated:   r.DateUpdated.Format(time.RFC3339),
		},
	}
}

type address struct {
	property.Address
}

func (a address) encode() *pb.Address {
	return &pb.Address{
		Line1: a.Line1,
		Line2: a.Line2,
		City:  a.City,
		State: a.State,
		Zip:   a.Zip,
	}
}

type rentalUnit struct {
	property.RentalUnit
}

func (u rentalUnit) encode() *pb.RentalUnit {
	return &pb.RentalUnit{
		Id:          u.ID.String(),
		Number:      u.Number,
		Bedrooms:    u.Bedrooms,
		Bathrooms:   u.Bathrooms,
		Size:        u.Size,
		RentAmount:  u.RentAmount,
		DateCreated: u.DateCreated.Format(time.RFC3339),
		DateUpdated: u.DateUpdated.Format(time.RFC3339),
	}
}
