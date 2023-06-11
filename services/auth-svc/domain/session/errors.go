package session

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrNotFound     = errors.New("session not found")
)
