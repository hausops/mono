package confirm

import (
	"errors"
)

var (
	ErrAlreadyConfirmed = errors.New("email already confirmed")
	ErrInvalidToken     = errors.New("invalid confirmation token")
	ErrNotFound         = errors.New("confirmation record not found")
)
