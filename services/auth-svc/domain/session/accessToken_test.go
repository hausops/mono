package session_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/session"
)

func TestParseAccessToken(t *testing.T) {
	t.Run("Invalid token", func(t *testing.T) {
		for _, tokenStr := range []string{
			"",
			"123",
			"ci6qsgrvq9l872j5bc8",
		} {
			_, err := session.ParseAccessToken(tokenStr)
			if err != session.ErrInvalidToken {
				t.Errorf("session.ParseAccessToken(%q) error = %v, want: ErrInvalidToken",
					tokenStr, err)
			}
		}
	})

	t.Run("Valid token", func(t *testing.T) {
		for _, tokenStr := range []string{
			"ci6qsgrvq9l872j5bc80",
			"00000000000000000000",
		} {
			token, err := session.ParseAccessToken(tokenStr)
			if err != nil {
				t.Errorf("session.ParseAccessToken(validToken) error = %v, want no error", err)
			}
			token2, _ := session.ParseAccessToken(token.String())
			if token != token2 {
				t.Error("Parsed access tokens from the same string are not equal")
			}
		}
	})
}

func TestAccessToken_String(t *testing.T) {
	// empty token
	{
		const emptyTokenStr = "00000000000000000000"

		var empty session.AccessToken
		{
			got := empty.String()
			if got != emptyTokenStr {
				t.Errorf("emptyToken.String() = %s, want %s", got, emptyTokenStr)
			}
		}

		parsedEmpty, err := session.ParseAccessToken(emptyTokenStr)
		if err != nil {
			t.Fatalf("session.ParseAccessToken(%s) err = %v", emptyTokenStr, err)
		}
		{
			got := parsedEmpty.String()
			if got != emptyTokenStr {
				t.Errorf("parsedEmpty.String() = %s, want %s", got, emptyTokenStr)
			}
		}

		// Literal creation is equivalent to parsing from the empty token string.
		if empty != parsedEmpty {
			t.Error("empty != parsedEmpty")
		}
	}

	// non-empty token
	{
		idStr := "cio035jjtoj2i2fbtpq0"
		id, err := session.ParseAccessToken(idStr)
		if err != nil {
			t.Fatalf("session.ParseAccessToken(%s) err = %v", idStr, err)
		}

		got := id.String()
		if got != idStr {
			t.Errorf("ID.String() = %s, want %s", got, idStr)
		}
	}
}
