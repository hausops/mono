package user

import (
	"net/mail"
	"time"
)

type User struct {
	ID          ID
	Email       mail.Address
	DateCreated time.Time
	DateUpdated time.Time
}
