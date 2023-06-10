package confirm

import (
	"errors"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidToken       = errors.New("invalid confirmation token")
	ErrPendingNotFound    = errors.New("pending confirmation not found")
)
