// Package confirm implements domain logic related to confirmation of
// user registration via email.
package confirm

import (
	"github.com/hausops/mono/services/user-svc/domain/user"
)

// Record represents whether an email is confirmed.
type Record struct {
	IsConfirmed bool
	Token       Token
	UserID      user.ID
}
