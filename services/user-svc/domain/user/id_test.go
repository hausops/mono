package user_test

import (
	"testing"

	"github.com/hausops/mono/services/user-svc/domain/user"
)

func TestNewID(t *testing.T) {
	id := user.NewID()
	if id.IsZero() {
		t.Error("user.NewID() generates an empty (zero) ID")
	}
}

func TestParseID(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		for _, idStr := range []string{
			"",
			"123",
			"ci6qsgrvq9l872j5bc8",
		} {
			_, err := user.ParseID(idStr)
			if err != user.ErrInvalidID {
				t.Errorf("user.ParseID(%q) error = %v, want: ErrInvalidID", idStr, err)
			}
		}
	})

	t.Run("Valid ID", func(t *testing.T) {
		for _, idStr := range []string{
			"ci6qsgrvq9l872j5bc80",
			"00000000000000000000",
		} {
			id, err := user.ParseID(idStr)
			if err != nil {
				t.Errorf("user.ParseID(%s) error = %v, want no error", idStr, err)
			}
			id2, _ := user.ParseID(id.String())
			if id != id2 {
				t.Error("Parsed ID from the same string are not equal")
			}
		}
	})
}

func TestID_IsZero(t *testing.T) {
	var empty user.ID
	if !empty.IsZero() {
		t.Error("emptyID.IsZero() = false, want true")
	}

	id := user.NewID()
	if id.IsZero() {
		t.Errorf("ID{%s}.IsZero() = true, want false", id)
	}
}

func TestID_String(t *testing.T) {
	// empty ID
	{
		const emptyIDStr = "00000000000000000000"

		var empty user.ID
		{
			got := empty.String()
			if got != emptyIDStr {
				t.Errorf("emptyID.String() = %s, want %s", got, emptyIDStr)
			}
		}

		parsedEmpty, err := user.ParseID(emptyIDStr)
		if err != nil {
			t.Fatalf("user.ParseID(%s) error = %v", emptyIDStr, err)
		}
		{
			got := parsedEmpty.String()
			if got != emptyIDStr {
				t.Errorf("parsedEmpty.String() = %s, want %s", got, emptyIDStr)
			}
		}

		// Literal creation is equivalent to parsing from the empty ID string.
		if empty != parsedEmpty {
			t.Error("empty != parsedEmpty")
		}
	}

	// non-empty ID
	{
		idStr := "cio035jjtoj2i2fbtpq0"
		id, err := user.ParseID(idStr)
		if err != nil {
			t.Fatalf("user.ParseID(%s) err = %v", idStr, err)
		}

		got := id.String()
		if got != idStr {
			t.Errorf("ID.String() = %s, want %s", got, idStr)
		}
	}
}
