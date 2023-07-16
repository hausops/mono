package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	repo user.Repository
}

func NewUserServer(repo user.Repository) *userServer {
	return &userServer{repo: repo}
}

func (s *userServer) Create(ctx context.Context, in *pb.EmailRequest) (*pb.User, error) {
	email, err := mail.ParseAddress(in.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	now := time.Now().UTC()
	u := user.User{
		ID:          user.NewID(),
		Email:       *email,
		DateCreated: now,
		DateUpdated: now,
	}

	created, err := s.repo.Upsert(ctx, u)
	if err != nil {
		if errors.Is(err, user.ErrEmailTaken) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, fmt.Errorf("user.Repository.Upsert(%s, %s): %w",
			u.ID, u.Email.Address, err)
	}

	return encodeUser(created), nil
}

func (s *userServer) FindByEmail(ctx context.Context, in *pb.EmailRequest) (*pb.User, error) {
	email, err := mail.ParseAddress(in.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	found, err := s.repo.FindByEmail(ctx, *email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, fmt.Errorf("user.Repository.FindByEmail(%s): %w", email.Address, err)
	}

	return encodeUser(found), nil
}

func encodeUser(u user.User) *pb.User {
	return &pb.User{
		Id:          u.ID.String(),
		Email:       u.Email.Address,
		DateCreated: u.DateCreated.Format(time.RFC3339),
		DateUpdated: u.DateUpdated.Format(time.RFC3339),
	}
}
