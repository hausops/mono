package confirm

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid confirmation token")
	ErrNotFound     = errors.New("confirmation record not found")
)
