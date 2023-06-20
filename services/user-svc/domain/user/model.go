package user

import (
	"net/mail"
	"time"

	"github.com/rs/xid"
)

type User struct {
	ID          xid.ID
	Email       mail.Address
	DateCreated time.Time
	DateUpdated time.Time
}
