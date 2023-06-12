package local

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/hausops/mono/services/auth-svc/domain/email"
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
  From: {{.From}}
  Subject: {{.Subject}}
  Body: {{.Body}}
`

func (ed *emailDispatcher) Send(_ context.Context, d email.Delivery, msg email.Message) error {
	body := msg.PlainText
	if msg.HTML != "" {
		body = msg.HTML
	}

	err := ed.tmpl.Execute(os.Stdout, struct {
		To      string
		From    string
		Subject string
		Body    string
	}{
		To:      d.To.String(),
		From:    d.From.String(),
		Subject: d.Subject,
		Body:    body,
	})
	if err != nil {
		return fmt.Errorf("cannot print email to Stdout: %w", err)
	}
	return nil
}
