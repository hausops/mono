package email

import (
	"context"
	"net/mail"
)

type Dispatcher interface {
	Send(ctx context.Context, to mail.Address, subject string, body string) error
}
