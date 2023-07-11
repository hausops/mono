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

	usr, err := s.user.Create(ctx, &userpb.EmailRequest{Email: email.Address})
	if err != nil {
		switch st := status.Convert(err); st.Code() {
		case codes.AlreadyExists:
			return credential.ErrAlreadyExists
		default:
			return fmt.Errorf("user.Create(%s): %w", email.Address, err)
		}
	}

	userID := usr.GetId()
	err = s.repos.Credential.Upsert(ctx, credential.Credential{
		Email:    email,
		Password: hashedPassword,
		UserID:   userID,
	})
	if err != nil {
		return fmt.Errorf("credential.Upsert(%s): %w", email.Address, err)
	}

	token := confirm.GenerateToken()
	err = s.sendConfirmEmail(ctx, email, token)
	if err != nil {
		return fmt.Errorf("sendConfirmEmail(%s): %w", email.Address, err)
	}

	err = s.repos.Confirm.Upsert(ctx, confirm.Record{
		IsConfirmed: false,
		Token:       token,
		UserID:      userID,
	})
	if err != nil {
		return fmt.Errorf("confirm.Upsert(%s): %w", userID, err)
	}
	return nil
}

func (s *Service) ResendConfirmationEmail(ctx context.Context, email mail.Address) error {
	cred, err := s.repos.Credential.FindByEmail(ctx, email)
	switch {
	case errors.Is(err, credential.ErrNotFound):
		return err
	case err != nil:
		return fmt.Errorf("credential.FindByEmail(%s): %w", email.Address, err)
	}

	rec, err := s.repos.Confirm.FindByUserID(ctx, cred.UserID)
	switch {
	case errors.Is(err, confirm.ErrNotFound):
		return err
	case err != nil:
		return fmt.Errorf("confirm.FindByUserID(%s): %w", cred.UserID, err)
	}

	if rec.IsConfirmed {
		return confirm.ErrAlreadyConfirmed
	}

	token := confirm.GenerateToken()
	err = s.sendConfirmEmail(ctx, email, token)
	if err != nil {
		return fmt.Errorf("sendConfirmEmail(%s): %w", email.Address, err)
	}

	err = s.repos.Confirm.Upsert(ctx, confirm.Record{
		IsConfirmed: false,
		Token:       token,
		UserID:      rec.UserID,
	})
	if err != nil {
		return fmt.Errorf("confirm.Upsert(%s): %w", rec.UserID, err)
	}
	return nil
}

func (s *Service) ConfirmEmail(ctx context.Context, token confirm.Token) (*session.Session, error) {
	rec, err := s.repos.Confirm.FindByToken(ctx, token)
	switch {
	case errors.Is(err, confirm.ErrNotFound):
		return nil, err
	case err != nil:
		return nil, fmt.Errorf("confirm.FindByToken(%s): %w", token, err)
	}

	if rec.IsConfirmed {
		return nil, confirm.ErrAlreadyConfirmed
	}

	err = s.repos.Confirm.Upsert(ctx, confirm.Record{
		IsConfirmed: true,
		UserID:      rec.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("confirm.Upsert(%s): %w", rec.UserID, err)
	}

	sess := session.New(rec.UserID, 24*time.Hour)
	err = s.repos.Session.Upsert(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("session.Upsert(%s): %w", rec.UserID, err)
	}
	return &sess, nil
}

func (s *Service) Login(ctx context.Context, email mail.Address, password []byte) (*session.Session, error) {
	cred, err := s.repos.Credential.FindByEmail(ctx, email)
	switch {
	case errors.Is(err, credential.ErrNotFound):
		return nil, err
	case err != nil:
		return nil, fmt.Errorf("credential.FindByEmail(%s): %w", email.Address, err)
	}

	err = bcrypt.CompareHashAndPassword(cred.Password, password)
	if err != nil {
		return nil, credential.ErrInvalidPassword
	}

	sess := session.New(cred.UserID, 24*time.Hour)
	err = s.repos.Session.Upsert(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("session.Upsert(%s): %w", cred.UserID, err)
	}

	return &sess, nil
}

func (s *Service) Logout(ctx context.Context, token session.AccessToken) error {
	err := s.repos.Session.DeleteByAccessToken(ctx, token)
	switch {
	case errors.Is(err, session.ErrNotFound):
		return err
	case err != nil:
		return fmt.Errorf("session.DeleteByAccessToken(%s): %w", token, err)
	}
	return nil
}

func (s *Service) CheckSession(ctx context.Context, token session.AccessToken) (*session.Session, error) {
	sess, err := s.repos.Session.FindByAccessToken(ctx, token)
	switch {
	case errors.Is(err, session.ErrNotFound):
		return nil, err
	case err != nil:
		return nil, fmt.Errorf("session.FindByAccessToken(%s): %w", token, err)
	}

	if sess.IsExpired() {
		go func() {
			err := s.repos.Session.DeleteByAccessToken(ctx, token)
			if err != nil {
				s.log.Error("Failed to delete expired session by access token",
					zap.String("accessToken", token.String()),
					zap.Error(err),
				)
			}
		}()
		return nil, session.ErrExpired
	}

	return &sess, nil
}

// sendConfirmEmail sends an email to the specified `to` address
// to confirm the email address.
func (s *Service) sendConfirmEmail(
	ctx context.Context,
	to mail.Address,
	token confirm.Token,
) error {
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

	return s.email.Send(ctx, delivery, msg)
}
