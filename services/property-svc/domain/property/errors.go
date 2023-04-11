package property

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID = errors.New("invalid property id")
	ErrNotFound  = errors.New("property not found")
)

type UpdateWrongPropertyTypeError struct {
	Property
	UpdateProperty
}

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
