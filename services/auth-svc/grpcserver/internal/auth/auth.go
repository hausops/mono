package auth

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/pb"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedAuthServer
	user       userpb.UserServiceClient
	credential *credential.Service
	confirm    *confirm.Service
}

func NewServer(
	user userpb.UserServiceClient,
	credential *credential.Service,
	confirm *confirm.Service,
) *server {
	return &server{
		user:       user,
		credential: credential,
		confirm:    confirm,
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

	err = s.credential.Save(ctx, *email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("save credential: %w", err)
	}

	err = s.confirm.SendEmail(ctx, *email)
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

	if _, err := s.credential.Lookup(ctx, *email); err != nil {
		return nil, status.Error(codes.NotFound, "Credential not found")
	}

	confirmed := s.confirm.IsConfirmed(ctx, *email)
	if confirmed {
		return nil, status.Error(codes.FailedPrecondition, "Email already confirmed")
	}

	err = s.confirm.SendEmail(ctx, *email)
	if err != nil {
		return nil, fmt.Errorf("send confirm email (%s): %w", email.Address, err)
	}

	return new(emptypb.Empty), nil
}
