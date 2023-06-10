// Package confirm implements domain logic related to confirmation of
// user registration via email.
package confirm

import (
	"errors"
	"net/mail"

	"github.com/rs/xid"
)

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

// generateToken creates a new randomized Token.
func generateToken() Token {
	return Token(xid.New())
}

func parseToken(raw []byte) (Token, error) {
	id, err := xid.FromBytes(raw)
	if errors.Is(err, xid.ErrInvalidID) {
		return Token{}, ErrInvalidToken
	}
	return Token(id), nil
}
