package user

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       mail.Address
	Name        string
	Verified    bool
	DateCreated time.Time
	DateUpdated time.Time
}
