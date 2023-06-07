// Package credential implements domain logic for managing user credentials.
//
// The package includes functionality to save and validate user credentials,
// using bcrypt for password hashing and comparison.
package credential

import "net/mail"

// Credential represents a user's email and hashed password.
type Credential struct {
	Email    mail.Address
	Password []byte
	// Expiration
}
