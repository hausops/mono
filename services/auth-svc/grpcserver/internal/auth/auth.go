package auth

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/domain/verification"
	"github.com/hausops/mono/services/auth-svc/pb"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedAuthServer
	user         userpb.UserServiceClient
	credential   *credential.Service
	verification *verification.Service
}

func NewServer(
	user userpb.UserServiceClient,
	credential *credential.Service,
	verification *verification.Service,
) *server {
	return &server{
		user:         user,
		credential:   credential,
		verification: verification,
	}
}

func (s *server) SignUp(ctx context.Context, r *pb.SignUpRequest) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	// TODO: password policy + strength
	if len(r.GetPassword()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}

	// TODO: rollback the user creation if any steps below fail
	_, err = s.user.Create(ctx, &userpb.EmailRequest{Email: email.Address})
	if err != nil {
		switch st, _ := status.FromError(err); st.Code() {
		case codes.AlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, fmt.Errorf("user.Create(%s): %w", email.Address, err)
		}
	}

	err = s.credential.Save(ctx, *email, r.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("save credential: %w", err)
	}

	err = s.verification.SendEmail(ctx, *email)
	if err != nil {
		return nil, fmt.Errorf("send verification email (%s): %w", email.Address, err)
	}

	return new(emptypb.Empty), nil
}
