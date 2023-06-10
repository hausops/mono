package email

import (
	"context"
	"errors"
	"net/mail"
)

type Dispatcher interface {
	Send(ctx context.Context, d Delivery, msg Message) error
}

type Delivery struct {
	To      mail.Address
	From    mail.Address
	Subject string
}

func (d Delivery) Validate() error {
	if d.To.Address == "" {
		return errors.New(`missing "to" address`)
	}
	if d.From.Address == "" {
		return errors.New(`missing "from" address`)
	}
	if d.Subject == "" {
		return errors.New("missing subject")
	}
	return nil
}

type Message struct {
	HTML      string
	PlainText string
}

func (m Message) Validate() error {
	if m.IsEmpty() {
		return errors.New("message is empty")
	}
	return nil
}

func (m Message) IsEmpty() bool {
	return m == Message{}
}
