package property

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/domain/property"
	"go.mongodb.org/mongo-driver/bson"
)

// This file defines local BSON models for reading and writing to mongo.

// propertyBSON is a BSON representation of the domain property.Property.
type propertyBSON interface {
	// isProperty is an empty receiver function to tag the concrete Property types.
	isProperty()
}

func toPropertyBSON(in property.Property) (propertyBSON, error) {
	switch t := in.(type) {
	case property.SingleFamilyProperty:
		return newSingleFamilyProperty(t), nil
	case property.MultiFamilyProperty:
		return newMultiFamilyProperty(t), nil
	default:
		return nil, property.UnhandledPropertyTypeError{Property: t}
	}
}

// decodePropertyFromBSON decodes a BSON raw document of a property to
// property.Property.
//
// The "_kind" field of the raw document determines the concrete property type
// to decode.
//
// Returns the decoded property and nil error if decoding was successful,
// otherwise returns nil and an error.
func decodePropertyFromBSON(raw bson.Raw) (property.Property, error) {
	s, ok := raw.Lookup("_kind").StringValueOK()
	if !ok {
		return nil, fmt.Errorf("missing _kind field in BSON document: %s", raw)
	}
	switch kind(s) {
	case kindSingleFamilyProperty:
		var sp singleFamilyProperty
		if err := bson.Unmarshal(raw, &sp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal BSON document: %w", err)
		}
		var p property.SingleFamilyProperty
		if err := sp.Decode(&p); err != nil {
			return nil, fmt.Errorf("failed to decode %s property from BSON document: %w", s, err)
		}
		return p, nil
	case kindMultiFamilyProperty:
		var mp multiFamilyProperty
		if err := bson.Unmarshal(raw, &mp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal BSON document: %w", err)
		}
		var p property.MultiFamilyProperty
		if err := mp.Decode(&p); err != nil {
			return nil, fmt.Errorf("failed to decode %s property from BSON document: %w", s, err)
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unknown property kind: %s", s)
	}
}

type singleFamilyProperty struct {
	ID            uuid.UUID  `bson:"_id"`
	Kind          kind       `bson:"_kind"`
	Address       address    `bson:"address"`
	CoverImageURL string     `bson:"cover_image_url,omitempty"`
	YearBuilt     int32      `bson:"year_built,omitempty"`
	Unit          rentalUnit `bson:"unit"`
	DateCreated   time.Time  `bson:"date_created"`
	DateUpdated   time.Time  `bson:"date_updated"`
}

func newSingleFamilyProperty(in property.SingleFamilyProperty) singleFamilyProperty {
	return singleFamilyProperty{
		ID:            in.ID,
		Kind:          kindSingleFamilyProperty,
		Address:       address(in.Address),
		CoverImageURL: in.CoverImageURL,
		YearBuilt:     in.YearBuilt,
		Unit:          rentalUnit(in.Unit),
		DateCreated:   in.DateCreated,
		DateUpdated:   in.DateUpdated,
	}
}

func (singleFamilyProperty) isProperty() {}

func (in singleFamilyProperty) Decode(p *property.SingleFamilyProperty) error {
	p.ID = in.ID
	p.Address = property.Address(in.Address)
	p.CoverImageURL = in.CoverImageURL
	p.YearBuilt = in.YearBuilt
	p.Unit = property.RentalUnit(in.Unit)
	p.DateCreated = in.DateCreated
	p.DateUpdated = in.DateUpdated
	return nil
}

type multiFamilyProperty struct {
	ID            uuid.UUID    `bson:"_id"`
	Kind          kind         `bson:"_kind"`
	Address       address      `bson:"address"`
	CoverImageURL string       `bson:"cover_image_url,omitempty"`
	YearBuilt     int32        `bson:"year_built,omitempty"`
	Units         []rentalUnit `bson:"units"`
	DateCreated   time.Time    `bson:"date_created"`
	DateUpdated   time.Time    `bson:"date_updated"`
}

func newMultiFamilyProperty(in property.MultiFamilyProperty) multiFamilyProperty {
	units := make([]rentalUnit, len(in.Units))
	for i, u := range in.Units {
		units[i] = rentalUnit(u)
	}

	return multiFamilyProperty{
		ID:            in.ID,
		Kind:          kindMultiFamilyProperty,
		Address:       address(in.Address),
		CoverImageURL: in.CoverImageURL,
		YearBuilt:     in.YearBuilt,
		Units:         units,
		DateCreated:   in.DateCreated,
		DateUpdated:   in.DateUpdated,
	}
}

func (multiFamilyProperty) isProperty() {}

func (in multiFamilyProperty) Decode(p *property.MultiFamilyProperty) error {
	units := make([]property.RentalUnit, len(in.Units))
	for i, u := range in.Units {
		units[i] = property.RentalUnit(u)
	}

	p.ID = in.ID
	p.Address = property.Address(in.Address)
	p.CoverImageURL = in.CoverImageURL
	p.YearBuilt = in.YearBuilt
	p.Units = units
	p.DateCreated = in.DateCreated
	p.DateUpdated = in.DateUpdated
	return nil
}

type kind string

const (
	kindSingleFamilyProperty kind = "single-family-property"
	kindMultiFamilyProperty  kind = "multi-family-property"
)

type address struct {
	Line1 string `bson:"line1"`
	Line2 string `bson:"line2,omitempty"`
	City  string `bson:"city"`
	State string `bson:"state"`
	Zip   string `bson:"zip"`
}

type rentalUnit struct {
	ID          uuid.UUID `bson:"_id"`
	Number      string    `bson:"number,omitempty"`
	Bedrooms    float32   `bson:"bedrooms,omitempty"`
	Bathrooms   float32   `bson:"bathrooms,omitempty"`
	Size        float32   `bson:"size,omitempty"`
	RentAmount  float32   `bson:"rent_amount,omitempty"`
	DateCreated time.Time `bson:"date_created"`
	DateUpdated time.Time `bson:"date_updated"`
}
