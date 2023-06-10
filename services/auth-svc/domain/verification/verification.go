// Package verification implements domain logic related to verifying
// a user's ownership of an email address.
package verification

import (
	"bytes"
	"net/mail"

	"github.com/rs/xid"
)

// Pending represents a pending email verification.
type Pending struct {
	Email mail.Address
	Token Token
}

// Token is a unique verification token.
type Token xid.ID

func (t Token) Equal(other Token) bool {
	return bytes.Equal(xid.ID(t).Bytes(), xid.ID(other).Bytes())
}

func (t Token) String() string {
	return xid.ID(t).String()
}

func GenerateToken() Token {
	return Token(xid.New())
}
