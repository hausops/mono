package verification

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/email"
)

type Service struct {
	pending  PendingRepository
	verified VerifiedRepository
	email    email.Dispatcher
}

func NewService(pending PendingRepository, verified VerifiedRepository, email email.Dispatcher) *Service {
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
	ver, err := s.pending.DeleteByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("delete pending verification: %w", err)
	}

	err = s.verified.Upsert(ctx, ver.Email)
	if err != nil {
		return fmt.Errorf("upsert verified email (%s): %w", ver.Email.Address, err)
	}
	return nil
}

// CheckVerified checks if the email address is verified.
// It returns true if the email address is verified, false otherwise.
// An error is returned if there's an issue retrieving the verification information.
func (s *Service) CheckVerified(ctx context.Context, email mail.Address) (bool, error) {
	err := s.verified.ExistByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrEmailNotVerified) {
			return false, nil
		}
		return false, fmt.Errorf("verified.ExistsByEmail(%s): %w", email.Address, err)
	}
	return true, nil
}
