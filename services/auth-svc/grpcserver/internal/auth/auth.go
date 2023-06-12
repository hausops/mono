package auth

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/auth"
	"github.com/hausops/mono/services/auth-svc/domain/confirm"
	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/domain/session"
	"github.com/hausops/mono/services/auth-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedAuthServer
	auth *auth.Service
}

func NewServer(auth *auth.Service) *server {
	return &server{auth: auth}
}

func (s *server) SignUp(ctx context.Context, r *pb.SignUpRequest) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email address")
	}

	err = s.auth.SignUp(ctx, *email, r.GetPassword())

	if err != nil {
		switch err {
		case credential.ErrAlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case credential.ErrInvalidPassword:
			return nil, status.Error(codes.InvalidArgument, "invalid password")
		}
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *server) ResendConfirmationEmail(ctx context.Context, r *pb.EmailRequest) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email address")
	}

	err = s.auth.ResendConfirmationEmail(ctx, *email)

	if err != nil {
		switch err {
		case credential.ErrNotFound, confirm.ErrNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case confirm.ErrAlreadyConfirmed:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *server) ConfirmEmail(ctx context.Context, r *pb.ConfirmEmailRequest) (*pb.Session, error) {
	token, err := confirm.ParseToken(r.GetToken())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token")
	}

	sess, err := s.auth.ConfirmEmail(ctx, token)

	if err != nil {
		switch err {
		case confirm.ErrNotFound:
			return nil, status.Error(codes.InvalidArgument, "invalid token")
		case confirm.ErrAlreadyConfirmed:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}
		return nil, err
	}

	res := &pb.Session{
		Email:       sess.Email.Address,
		AccessToken: sess.AccessToken.String(),
	}
	return res, nil
}

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.Session, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email address")
	}

	sess, err := s.auth.Login(ctx, *email, r.GetPassword())

	if err != nil {
		switch err {
		case credential.ErrNotFound:
			// Return as "not found" from service back-end.
			// It will be turned to permission denied: invalid credential in auth-api.
			return nil, status.Error(codes.NotFound, err.Error())
		case credential.ErrInvalidPassword:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	res := &pb.Session{
		Email:       sess.Email.Address,
		AccessToken: sess.AccessToken.String(),
	}
	return res, nil
}

func (s *server) Logout(ctx context.Context, r *pb.LogoutRequest) (*emptypb.Empty, error) {
	accessToken, err := session.ParseAccessToken(r.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token")
	}

	err = s.auth.Logout(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *server) CheckSession(ctx context.Context, r *pb.CheckSessionRequest) (*pb.CheckSessionResponse, error) {
	accessToken, err := session.ParseAccessToken(r.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token")
	}

	sess, err := s.auth.CheckSession(ctx, accessToken)

	if err != nil {
		switch err {
		case session.ErrExpired, session.ErrNotFound:
			return &pb.CheckSessionResponse{Valid: false}, nil
		}
		return nil, err
	}

	res := &pb.CheckSessionResponse{
		Valid: true,
		Email: sess.Email.Address,
	}
	return res, nil
}
