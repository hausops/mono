package property

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

// UnhandledPropertyTypeError is meant for using in type switch on Property
// to return when a concrete type is unexpected.
type UnhandledPropertyTypeError struct {
	Property Property
}

func (e *UnhandledPropertyTypeError) Error() string {
	return fmt.Sprintf("unhandled Property type: %T", e.Property)
}
