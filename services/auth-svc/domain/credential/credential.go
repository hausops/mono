// Package credential ...
package credential

import "net/mail"

type Credential struct {
	Email    mail.Address
	Password []byte
	// Expiration
}
