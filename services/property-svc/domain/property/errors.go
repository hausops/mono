package property

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound  = errors.New("property not found")
	ErrInvalidID = errors.New("invalid property id")
)

// UnhandledPropertyTypeError is intended for use in a type switch on Property
// to handle cases when an unexpected concrete type is encountered.
type UnhandledPropertyTypeError struct {
	Property Property
}

// Error implements the error interface for UnhandledPropertyTypeError.
func (e UnhandledPropertyTypeError) Error() string {
	return fmt.Sprintf("unhandled Property type: %T", e.Property)
}
