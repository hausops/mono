// Package verification implements domain logic related to user's
// proof of ownership (verification) over an email address.
package verification

import (
	"bytes"
	"net/mail"

	"github.com/rs/xid"
)

// PendingVerification represents ...
type PendingVerification struct {
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
