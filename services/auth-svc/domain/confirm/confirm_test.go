package confirm_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

func TestParseToken(t *testing.T) {
	t.Run("invalid token", func(t *testing.T) {
		for _, tokenStr := range []string{"", "ci6qsgrvq9l872j5bc8"} {
			_, err := confirm.ParseToken(tokenStr)
			if err != confirm.ErrInvalidToken {
				t.Errorf("confirm.ParseToken(%q) error = %v, want: ErrInvalidToken",
					tokenStr, err)
			}
		}
	})

	t.Run("valid token", func(t *testing.T) {
		token, err := confirm.ParseToken("ci6qsgrvq9l872j5bc80")
		if err != nil {
			t.Errorf("confirm.ParseToken(validToken) error = %v, want no error", err)
		}
		token2, _ := confirm.ParseToken(token.String())
		if token != token2 {
			t.Error("Parsed access tokens from the same string are not equal")
		}
	})
}
