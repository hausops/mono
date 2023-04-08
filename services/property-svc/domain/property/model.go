package property

import "github.com/google/uuid"

type Property interface {
	// isProperty is an empty receiver function to tag the concrete Property types.
	isProperty()
}

type SingleFamilyProperty struct {
	ID            uuid.UUID
	Address       Address
	CoverImageURL string
	YearBuilt     int32
	Unit          RentalUnit
}

func (p SingleFamilyProperty) isProperty() {}

type MultiFamilyProperty struct {
	ID            uuid.UUID
	Address       Address
	CoverImageURL string
	YearBuilt     int32
	Units         []RentalUnit
}

func (p MultiFamilyProperty) isProperty() {}

type Address struct {
	Line1 string
	Line2 string
	City  string
	State string
	Zip   string
}

type RentalUnit struct {
	ID         uuid.UUID
	Number     string
	Bedrooms   float32
	Bathrooms  float32
	Size       float32
	RentAmount float32
	// ActiveListing
}
