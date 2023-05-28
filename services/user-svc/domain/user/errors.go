package user

import (
	"errors"
)

var (
	ErrEmailTaken = errors.New("email already taken")
	ErrInvalidID  = errors.New("invalid user id")
	ErrMissingID  = errors.New("missing user id")
	ErrNotFound   = errors.New("user not found")
)
