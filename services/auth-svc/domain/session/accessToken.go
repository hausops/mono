package session

import "github.com/rs/xid"

type AccessToken xid.ID

func NewAccessToken() AccessToken {
	return AccessToken(xid.New())
}

func (at AccessToken) String() string {
	return xid.ID(at).String()
}

func ParseAccessToken(s string) (AccessToken, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return AccessToken{}, ErrInvalidToken
	}
	return AccessToken(id), nil
}
