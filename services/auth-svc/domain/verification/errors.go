package verification

import (
	"errors"
)

var (
	ErrEmailNotVerified = errors.New("the email is not verified")
	ErrPendingNotFound  = errors.New("pending verification record not found")
)
