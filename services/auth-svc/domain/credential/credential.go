// Package credential implements domain logic for managing user credentials.
package credential

import (
	"net/mail"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

// Credential represents a user's email and hashed password.
type Credential struct {
	Email    mail.Address
	Password []byte
}

const minPasswordEntropy = 50

// ValidatePassword ensures the password conforms to our password policy.
func ValidatePassword(password string) error {
	return passwordvalidator.Validate(password, minPasswordEntropy)
}

// HashPassword hashes the password to save to the repo.
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
