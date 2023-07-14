package user

import "github.com/rs/xid"

type ID xid.ID

// NewID creates a new non-empty ID.
func NewID() ID {
	return ID(xid.New())
}

// ParseID parse s to ID or an error if s is not a valid ID.
func ParseID(s string) (ID, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return ID{}, ErrInvalidID
	}
	return ID(id), nil
}

func (id ID) IsZero() bool {
	return xid.ID(id).IsZero()
}

func (id ID) String() string {
	return xid.ID(id).String()
}
