package auth

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/domain/email"
	"github.com/hausops/mono/services/auth-svc/domain/session"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	user  userpb.UserServiceClient
	repos Repositories
	email email.Dispatcher
	log   *zap.Logger
}

// Repositories groups the repositories needed by Service.
type Repositories struct {
	Confirm    confirm.Repository
	Credential credential.Repository
	Session    session.Repository
}

func NewService(
	user userpb.UserServiceClient,
	repos Repositories,
	email email.Dispatcher,
	log *zap.Logger,
) *Service {
	return &Service{
		user:  user,
		repos: repos,
		email: email,
		log:   log,
	}
}

func (s *Service) SignUp(ctx context.Context, email mail.Address, password []byte) error {
	err := credential.ValidatePassword(string(password))
	if err != nil {
		return credential.ErrInvalidPassword
	}

	hashedPassword, err := credential.HashPassword(password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	_, err = s.user.Create(ctx, &userpb.EmailRequest{Email: email.Address})
	if err != nil {
		switch st := status.Convert(err); st.Code() {
		case codes.AlreadyExists:
			return credential.ErrAlreadyExists
		default:
			return fmt.Errorf("user.Create(%s): %w", email.Address, err)
		}
	}

	cred := credential.Credential{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.repos.Credential.Upsert(ctx, cred)
	if err != nil {
		return fmt.Errorf("credential.Upsert(%s): %w", email.Address, err)
	}

	err = s.sendConfirmEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("sendConfirmEmail(%s): %w", email.Address, err)
	}

	return nil
}

func (s *Service) ResendConfirmationEmail(ctx context.Context, email mail.Address) error {
	if _, err := s.repos.Credential.FindByEmail(ctx, email); err != nil {
		return err
	}

	rec, err := s.repos.Confirm.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if rec.IsConfirmed {
		return confirm.ErrAlreadyConfirmed
	}

	err = s.sendConfirmEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("sendConfirmEmail(%s): %w", email.Address, err)
	}

	return nil
}

func (s *Service) ConfirmEmail(ctx context.Context, token confirm.Token) (*session.Session, error) {
	rec, err := s.repos.Confirm.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if rec.IsConfirmed {
		return nil, confirm.ErrAlreadyConfirmed
	}

	confirmed := confirm.Record{
		Email:       rec.Email,
		IsConfirmed: true,
		Token:       confirm.Token{},
	}

	err = s.repos.Confirm.Upsert(ctx, confirmed)
	if err != nil {
		return nil, fmt.Errorf("confirmRepo.Upsert(%s): %w", rec.Email.Address, err)
	}

	sess := session.New(confirmed.Email, 24*time.Hour)
	err = s.repos.Session.Upsert(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("save session for %s: %w", sess.Email.Address, err)
	}

	return &sess, nil
}

func (s *Service) Login(ctx context.Context, email mail.Address, password []byte) (*session.Session, error) {
	cred, err := s.repos.Credential.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, credential.ErrNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("credential.FindByEmail(%s): %w", email.Address, err)
	}

	err = bcrypt.CompareHashAndPassword(cred.Password, password)
	if err != nil {
		return nil, credential.ErrInvalidPassword
	}

	sess := session.New(cred.Email, 24*time.Hour)
	err = s.repos.Session.Upsert(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("save session for %s: %w", sess.Email.Address, err)
	}

	return &sess, nil
}

func (s *Service) Logout(ctx context.Context, token session.AccessToken) error {
	sess, err := s.repos.Session.FindByAccessToken(ctx, token)
	if err != nil {
		return fmt.Errorf("find session by access token (%s): %w", token, err)
	}

	_, err = s.repos.Session.DeleteByEmail(ctx, sess.Email)
	if err != nil {
		return fmt.Errorf("delete session by email (%s): %w", sess.Email.Address, err)
	}
	return nil
}

func (s *Service) CheckSession(ctx context.Context, token session.AccessToken) (*session.Session, error) {
	sess, err := s.repos.Session.FindByAccessToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if sess.IsExpired() {
		go func() {
			_, err := s.repos.Session.DeleteByEmail(ctx, sess.Email)
			if err != nil {
				s.log.Error("Failed to delete expired session by email",
					zap.String("email", sess.Email.Address),
					zap.Error(err),
				)
			}
		}()
		return nil, session.ErrExpired
	}

	return &sess, nil
}

// sendConfirmEmail generates a new confirm record and sends an email to
// the specified `to` address to confirm the email address.
func (s *Service) sendConfirmEmail(ctx context.Context, to mail.Address) error {
	token := confirm.GenerateToken()

	{
		delivery := email.Delivery{
			To:      to,
			From:    mail.Address{Name: "HausOps", Address: "no-reply@hausops.com"},
			Subject: "Confirm your email address to start using HausOps",
		}

		msg := email.Message{
			PlainText: fmt.Sprintf(
				"Confirm your email address: https://auth.hausops.com/confirm?t=%s",
				token,
			),
		}

		err := s.email.Send(ctx, delivery, msg)
		if err != nil {
			return err
		}
	}

	rec := confirm.Record{
		Email:       to,
		Token:       token,
		IsConfirmed: false,
	}

	return s.repos.Confirm.Upsert(ctx, rec)
}
