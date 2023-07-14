// Package confirm implements domain logic related to confirmation of
// user registration via email.
package confirm

import (
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/rs/xid"
)

// Record represents whether an email is confirmed.
type Record struct {
	IsConfirmed bool
	Token       Token
	UserID      user.ID
}

// Token is a unique token for email confirmation.
type Token xid.ID

func (t Token) IsZero() bool {
	return xid.ID(t).IsZero()
}

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
