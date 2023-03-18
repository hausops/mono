package property

import "fmt"

type NotFoundError struct {
	// ID is the ID of the property for the lookup
	ID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("property with id=%q is not found.", e.ID)
}
