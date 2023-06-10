package local

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"text/template"
)

type emailDispatcher struct {
	tmpl *template.Template
}

func NewEmailDispatcher() *emailDispatcher {
	return &emailDispatcher{
		tmpl: template.Must(template.New("local").Parse(emailTemplate)),
	}
}

const emailTemplate = `
Send email
  To: {{.To}}
  Subject: {{.Subject}}
  Body: {{.Body}}
`

func (ed *emailDispatcher) Send(
	_ context.Context, to mail.Address, subject string, body string,
) error {
	err := ed.tmpl.Execute(os.Stdout, struct {
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
