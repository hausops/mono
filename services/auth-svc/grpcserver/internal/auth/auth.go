package auth

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/domain/email"
	"github.com/hausops/mono/services/auth-svc/pb"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedAuthServer
	user           userpb.UserServiceClient
	credentialRepo credential.Repository
	confirmRepo    confirm.Repository
	email          email.Dispatcher
}

func NewServer(
	user userpb.UserServiceClient,
	credentialRepo credential.Repository,
	confirmRepo confirm.Repository,
	email email.Dispatcher,
) *server {
	return &server{
		user:           user,
		credentialRepo: credentialRepo,
		confirmRepo:    confirmRepo,
		email:          email,
	}
}

func (s *server) SignUp(ctx context.Context, r *pb.SignUpRequest) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	password := r.GetPassword()
	err = credential.ValidatePassword(string(password))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}

	hashedPassword, err := credential.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	_, err = s.user.Create(ctx, &userpb.EmailRequest{Email: email.Address})
	if err != nil {
		switch st, _ := status.FromError(err); st.Code() {
		case codes.AlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, fmt.Errorf("user.Create(%s): %w", email.Address, err)
		}
	}

	cred := credential.Credential{
		Email:    *email,
		Password: hashedPassword,
	}

	err = s.credentialRepo.Upsert(ctx, cred)
	if err != nil {
		return nil, fmt.Errorf("credentialRepo.Upsert(%s): %w", email.Address, err)
	}

	err = s.sendConfirmEmail(ctx, *email)
	if err != nil {
		return nil, fmt.Errorf("send confirm email (%s): %w", email.Address, err)
	}

	return new(emptypb.Empty), nil
}

func (s *server) ResendConfirmationEmail(
	ctx context.Context,
	r *pb.ResendConfirmationEmailRequest,
) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	if _, err := s.credentialRepo.FindByEmail(ctx, *email); err != nil {
		return nil, status.Error(codes.NotFound, "credential not found")
	}

	rec, err := s.confirmRepo.FindByEmail(ctx, *email)
	if err != nil {
		return nil, err
	}

	if rec.IsConfirmed {
		return nil, status.Error(codes.FailedPrecondition, "Email already confirmed")
	}

	err = s.sendConfirmEmail(ctx, *email)
	if err != nil {
		return nil, fmt.Errorf("send confirm email (%s): %w", email.Address, err)
	}

	return new(emptypb.Empty), nil
}

// sendConfirmEmail generates a new confirm record and sends an email to
// the specified `to` address to confirm the email address.
func (s *server) sendConfirmEmail(ctx context.Context, to mail.Address) error {
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
		Token:       &token,
		IsConfirmed: false,
	}

	return s.confirmRepo.Upsert(ctx, rec)
}
