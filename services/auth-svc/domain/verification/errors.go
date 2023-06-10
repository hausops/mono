package verification

import (
	"errors"
)

var (
	// pending
	ErrPendingNotFound = errors.New("pending verification record not found")

	// verified email
	ErrEmailAlreadyExists = errors.New("email already exists")
)
