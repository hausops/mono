package property

import (
	"time"

	"github.com/google/uuid"
)

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
	DateCreated   time.Time
	DateUpdated   time.Time
}

func (p SingleFamilyProperty) isProperty() {}

type MultiFamilyProperty struct {
	ID            uuid.UUID
	Address       Address
	CoverImageURL string
	YearBuilt     int32
	Units         []RentalUnit
	DateCreated   time.Time
	DateUpdated   time.Time
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
	ID          uuid.UUID
	Number      string
	Bedrooms    float32
	Bathrooms   float32
	Size        float32
	RentAmount  float32
	DateCreated time.Time
	DateUpdated time.Time
	// ActiveListing
}

type UpdateProperty interface {
	// isUpdateProperty is an empty receiver function to tag the concrete Property types.
	isUpdateProperty()
}

type UpdateSingleFamilyProperty struct {
	Address       *UpdateAddress
	CoverImageURL *string
	YearBuilt     *int32
	Unit          *UpdateRentalUnit
}

func (up UpdateSingleFamilyProperty) isUpdateProperty() {}

type UpdateMultiFamilyProperty struct {
	Address       *UpdateAddress
	CoverImageURL *string
	YearBuilt     *int32
}

func (up UpdateMultiFamilyProperty) isUpdateProperty() {}

type UpdateAddress struct {
	Line1 *string
	Line2 *string
	City  *string
	State *string
	Zip   *string
}

type UpdateRentalUnit struct {
	Number     *string
	Bedrooms   *float32
	Bathrooms  *float32
	Size       *float32
	RentAmount *float32
}
