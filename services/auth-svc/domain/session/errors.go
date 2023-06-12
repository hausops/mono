package session

import "errors"

var (
	ErrExpired      = errors.New("session expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrNotFound     = errors.New("session not found")
)
