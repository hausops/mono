package property

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID = errors.New("invalid property id")
	ErrNotFound  = errors.New("property not found")
)

// MissingIDError is returned when an ID is missing.
type MissingIDError struct {
	// Message is a human-readable description of the error.
	Message string
}

// Error implements the error interface for MissingIDError.
func (e MissingIDError) Error() string {
	return e.Message
}

// UpdateWrongPropertyTypeError is returned when attempting to update
// a property with an incompatible input type.
type UpdateWrongPropertyTypeError struct {
	// Property is the property that was being updated.
	Property

	// UpdateProperty is the input used to update the property.
	UpdateProperty
}

// Error implements the error interface for UpdateWrongPropertyTypeError.
func (e UpdateWrongPropertyTypeError) Error() string {
	return fmt.Sprintf("wrong property type: updating %T with %T", e.Property, e.UpdateProperty)
}

// UnhandledPropertyTypeError is intended for use in a type switch on Property
// to handle cases when an unexpected concrete type is encountered.
type UnhandledPropertyTypeError struct {
	Property
}

// Error implements the error interface for UnhandledPropertyTypeError.
func (e UnhandledPropertyTypeError) Error() string {
	return fmt.Sprintf("unhandled Property type[%T]", e.Property)
}
