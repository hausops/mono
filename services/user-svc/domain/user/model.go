package user

import (
	"net/mail"
	"time"

	// "github.com/google/uuid"
	"github.com/speps/go-hashids/v2"
)

type User struct {
	// ID          uuid.UUID
	ID          hashids.HashID
	Email       mail.Address
	Name        string
	Verified    bool
	DateCreated time.Time
	DateUpdated time.Time
}
