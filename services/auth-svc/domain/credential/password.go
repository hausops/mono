package credential

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordEntropy = 50

// ValidatePassword ensures the password conforms to our password policy.
func ValidatePassword(password string) error {
	return passwordvalidator.Validate(password, minPasswordEntropy)
}

// HashPassword hashes the password to save to the repo.
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
