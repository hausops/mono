package credential

import (
	"errors"
)

var (
	ErrAlreadyExists   = errors.New("credential already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrNotFound        = errors.New("credential not found")
)
