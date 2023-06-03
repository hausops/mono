package local

import (
	"context"
	"fmt"
	"net/mail"
)

type emailDispatcher struct{}

func NewEmailDispatcher() *emailDispatcher {
	return &emailDispatcher{}
}

func (ed *emailDispatcher) Send(
	_ context.Context, to mail.Address, subject string, body string,
) error {

	const template = `
Send email
	To: %s
	Subject: %s
	Body: %s`

	_, err := fmt.Printf(template, to.Address, subject, body)
	if err != nil {
		return fmt.Errorf("cannot print email to Stdout: %w", err)
	}
	return nil
}
