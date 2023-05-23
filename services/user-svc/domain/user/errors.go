package user

import (
	"errors"
)

var (
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrInvalidID        = errors.New("invalid user id")
	ErrNotFound         = errors.New("user not found")
)
