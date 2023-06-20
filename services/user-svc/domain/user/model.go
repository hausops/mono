package user

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       mail.Address
	DateCreated time.Time
	DateUpdated time.Time
}
