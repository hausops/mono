package user

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	repo user.Repository
	// svc  *user.Service
}

func NewServer(repo user.Repository) *server {
	return &server{repo: repo}
}

func (s *server) Create(ctx context.Context, in *pb.EmailRequest) (*pb.User, error) {
	emailPtr, err := mail.ParseAddress(in.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	now := time.Now().UTC()
	u := user.User{
		ID:          uuid.New(),
		Email:       *emailPtr,
		Verified:    false,
		DateCreated: now,
		DateUpdated: now,
	}

	created, err := s.repo.Upsert(ctx, u)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, fmt.Errorf("user.Repository.Upsert(%s, %s): %w",
			u.ID, u.Email.Address, err)
	}

	return encodeUser(created), nil
}

func (s *server) FindByEmail(ctx context.Context, in *pb.EmailRequest) (*pb.User, error) {
	emailPtr, err := mail.ParseAddress(in.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	email := *emailPtr
	found, err := s.repo.FindByEmail(ctx, email)
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
		Name:        u.Name,
		Verified:    u.Verified,
		DateCreated: u.DateCreated.Format(time.RFC3339),
		DateUpdated: u.DateUpdated.Format(time.RFC3339),
	}
}
