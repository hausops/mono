package credential

import (
	"errors"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrNotFound        = errors.New("credential not found")
)
