package confirm

import "github.com/rs/xid"

// Token is a unique token for email confirmation.
type Token xid.ID

// NewToken creates a new randomized Token.
func NewToken() Token {
	return Token(xid.New())
}

func ParseToken(s string) (Token, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return Token{}, ErrInvalidToken
	}
	return Token(id), nil
}

func (t Token) IsZero() bool {
	return xid.ID(t).IsZero()
}

func (t Token) String() string {
	return xid.ID(t).String()
}
