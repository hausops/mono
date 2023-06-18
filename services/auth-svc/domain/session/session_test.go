package session_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/session"
)

func TestParseToken(t *testing.T) {
	t.Run("invalid token", func(t *testing.T) {
		for _, tokenStr := range []string{"", "ci6qsgrvq9l872j5bc8"} {
			_, err := session.ParseAccessToken(tokenStr)
			if err != session.ErrInvalidToken {
				t.Errorf("session.ParseAccessToken(%q) error = %v, want: ErrInvalidToken",
					tokenStr, err)
			}
		}
	})

	t.Run("valid token", func(t *testing.T) {
		for _, tokenStr := range []string{
			"ci6qsgrvq9l872j5bc80",
			"ci6s2jjvq9l9rjp3b390",
		} {
			token, err := session.ParseAccessToken(tokenStr)
			if err != nil {
				t.Errorf("session.ParseAccessToken(validToken) error = %v, want no error", err)
			}
			token2, _ := session.ParseAccessToken(tokenStr)
			if token != token2 {
				t.Error("Parsed access tokens from the same string are not equal")
			}
		}
	})
}
