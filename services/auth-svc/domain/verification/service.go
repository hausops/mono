package verification

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/email"
)

type Service struct {
	pending  PendingRepository
	verified VerifiedEmailRepository
	email    email.Dispatcher
}

func NewService(pending PendingRepository, verified VerifiedEmailRepository, email email.Dispatcher) *Service {
	return &Service{
		pending:  pending,
		verified: verified,
		email:    email,
	}
}

// SendEmail generates a new verification token and sends an email to
// the specified `to` address in order to verify the email address.
func (s *Service) SendEmail(ctx context.Context, to mail.Address) error {
	token := generateToken()

	subject := "Verify your email to start using HausOps"
	body := fmt.Sprintf("Verify your email address: https://auth.hausops.com/verify-email?t=%s", token)
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

// Verify marks the email address associated with the token as verified
// and deletes the corresponding pending verification record for the token.
func (s *Service) Verify(ctx context.Context, token Token) error {
	pending, err := s.pending.DeleteByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("delete pending verification: %w", err)
	}

	email := pending.Email
	err = s.verified.Add(ctx, email)
	if err != nil {
		return fmt.Errorf("upsert verified email (%s): %w", email.Address, err)
	}
	return nil
}

// IsVerified checks if the email address is verified.
// It returns true if the email address is verified, false otherwise.
// An error is returned if there's an issue retrieving the verification information.
func (s *Service) IsVerified(ctx context.Context, email mail.Address) bool {
	return s.verified.Exist(ctx, email)
}
