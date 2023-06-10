package confirm

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/email"
)

type Service struct {
	pending   PendingRepository
	confirmed ConfirmedEmailRepository
	email     email.Dispatcher
}

func NewService(
	pending PendingRepository,
	confirmed ConfirmedEmailRepository,
	email email.Dispatcher,
) *Service {
	return &Service{
		pending:   pending,
		confirmed: confirmed,
		email:     email,
	}
}

// SendEmail generates a new pending confirmation and sends an email to
// the specified `to` address to confirm the email address.
func (s *Service) SendEmail(ctx context.Context, to mail.Address) error {
	token := generateToken()

	subject := "Confirm your email address to start using HausOps"
	body := fmt.Sprintf("Confirm your email address: https://auth.hausops.com/confirm?t=%s", token)
	err := s.email.Send(ctx, to, subject, body)
	if err != nil {
		return err
	}

	err = s.pending.Upsert(ctx, Pending{Email: to, Token: token})
	if err != nil {
		return err
	}
	return nil
}

// Confirm marks the email address associated with the token as confirmed
// removing the pending confirmation.
//
// If the pending confirmation for the token does not exist,
// ErrPendingNotFound is returned.
//
// If the email is already confirmed, ErrEmailAlreadyExists is returned.
func (s *Service) Confirm(ctx context.Context, rawToken []byte) error {
	token, err := parseToken(rawToken)
	if err != nil {
		return err
	}

	pending, err := s.pending.DeleteByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("delete pending confirmation: %w", err)
	}

	email := pending.Email
	err = s.confirmed.Add(ctx, email)
	if err != nil {
		return fmt.Errorf("upsert confirmed email (%s): %w", email.Address, err)
	}
	return nil
}

// IsConfirmed checks if the email address is confirmed.
func (s *Service) IsConfirmed(ctx context.Context, email mail.Address) bool {
	return s.confirmed.Exist(ctx, email)
}
