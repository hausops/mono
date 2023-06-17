// Package confirm implements domain logic related to confirmation of
// user registration via email.
package confirm

import (
	"net/mail"

	"github.com/rs/xid"
)

// Record represents whether an email is confirmed.
type Record struct {
	Email       mail.Address
	IsConfirmed bool
	Token       *Token
}

// Pending represents a pending email confirmation.
type Pending struct {
	Email mail.Address
	Token Token
}

// Token is a unique token for email confirmation.
type Token xid.ID

func (t Token) String() string {
	return xid.ID(t).String()
}

// GenerateToken creates a new randomized Token.
func GenerateToken() Token {
	return Token(xid.New())
}

func ParseToken(s string) (Token, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return Token{}, ErrInvalidToken
	}
	return Token(id), nil
}
