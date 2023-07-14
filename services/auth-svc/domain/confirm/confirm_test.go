package confirm_test

import (
	"testing"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/rs/xid"
)

func TestParseToken(t *testing.T) {
	t.Run("Invalid token", func(t *testing.T) {
		for _, tokenStr := range []string{"", "ci6qsgrvq9l872j5bc8"} {
			_, err := confirm.ParseToken(tokenStr)
			if err != confirm.ErrInvalidToken {
				t.Errorf("confirm.ParseToken(%q) error = %v, want: ErrInvalidToken",
					tokenStr, err)
			}
		}
	})

	t.Run("Valid token", func(t *testing.T) {
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

func TestToken_IsZero(t *testing.T) {
	var empty confirm.Token
	if !empty.IsZero() {
		t.Error("emptyToken.IsZero() = false, want true")
	}

	token := confirm.GenerateToken()
	if token.IsZero() {
		t.Errorf("Token{%s}.IsZero() = true, want false", token)
	}
}

func TestToken_String(t *testing.T) {
	// empty token
	{
		var empty confirm.Token
		got := empty.String()
		want := "00000000000000000000"
		if got != want {
			t.Errorf("emptyToken.String() = %s, want %s", got, want)
		}
	}

	// non-empty token
	{
		tokenStr := "cio035jjtoj2i2fbtpq0"
		id, err := xid.FromString(tokenStr)
		if err != nil {
			t.Fatalf("xid.FromString(%s) err = %v", tokenStr, err)
		}
		token := confirm.Token(id)

		got := token.String()
		want := tokenStr
		if got != want {
			t.Errorf("token.String() = %s, want %s", got, want)
		}
	}
}
