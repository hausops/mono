package local

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"text/template"
)

type emailDispatcher struct{}

func NewEmailDispatcher() *emailDispatcher {
	return &emailDispatcher{}
}

func (ed *emailDispatcher) Send(
	_ context.Context, to mail.Address, subject string, body string,
) error {
	tmpl := template.Must(template.ParseFiles("adapter/local/email.txt"))
	err := tmpl.Execute(os.Stdout, struct {
		To      string
		Subject string
		Body    string
	}{
		To:      to.Address,
		Subject: subject,
		Body:    body,
	})
	if err != nil {
		return fmt.Errorf("cannot print email to Stdout: %w", err)
	}
	return nil
}
