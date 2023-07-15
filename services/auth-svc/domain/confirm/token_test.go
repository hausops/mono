package confirm_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
)

func TestGenerateToken(t *testing.T) {
	token := confirm.NewToken()
	if token.IsZero() {
		t.Error("confirm.GenerateToken() generates an empty (zero) token")
	}
}

func TestParseToken(t *testing.T) {
	t.Run("Invalid token", func(t *testing.T) {
		for _, tokenStr := range []string{
			"",
			"123",
			"ci6qsgrvq9l872j5bc8",
		} {
			_, err := confirm.ParseToken(tokenStr)
			if err != confirm.ErrInvalidToken {
				t.Errorf("confirm.ParseToken(%q) error = %v, want: ErrInvalidToken",
					tokenStr, err)
			}
		}
	})

	t.Run("Valid token", func(t *testing.T) {
		for _, tokenStr := range []string{
			"ci6qsgrvq9l872j5bc80",
			"00000000000000000000",
		} {
			token, err := confirm.ParseToken(tokenStr)
			if err != nil {
				t.Errorf("confirm.ParseToken(validToken) error = %v, want no error", err)
			}
			token2, _ := confirm.ParseToken(token.String())
			if token != token2 {
				t.Error("Parsed access tokens from the same string are not equal")
			}
		}
	})
}

func TestToken_IsZero(t *testing.T) {
	var empty confirm.Token
	if !empty.IsZero() {
		t.Error("emptyToken.IsZero() = false, want true")
	}

	token := confirm.NewToken()
	if token.IsZero() {
		t.Errorf("Token{%s}.IsZero() = true, want false", token)
	}
}

func TestToken_String(t *testing.T) {
	// empty token
	{
		const emptyTokenStr = "00000000000000000000"

		var empty confirm.Token
		{
			got := empty.String()
			if got != emptyTokenStr {
				t.Errorf("emptyToken.String() = %s, want %s", got, emptyTokenStr)
			}
		}

		parsedEmpty, err := confirm.ParseToken(emptyTokenStr)
		if err != nil {
			t.Fatalf("confirm.ParseToken(%s) err = %v", emptyTokenStr, err)
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
		id, err := confirm.ParseToken(idStr)
		if err != nil {
			t.Fatalf("confirm.ParseToken(%s) err = %v", idStr, err)
		}

		got := id.String()
		if got != idStr {
			t.Errorf("ID.String() = %s, want %s", got, idStr)
		}
	}
}
