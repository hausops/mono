package user

import (
	"context"
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

	if _, err := s.repo.Upsert(ctx, u); err != nil {
		return nil, fmt.Errorf("user.Repository.Upsert(%v): %w", u, err)
	}
	return fromUser(u), nil
}

func (s *server) FindByEmail(ctx context.Context, in *pb.EmailRequest) (*pb.User, error) {
	emailPtr, err := mail.ParseAddress(in.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	email := *emailPtr
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user.Repository.FindByEmail(%v): %w", email.Address, err)
	}
	return fromUser(u), nil
}

func fromUser(u user.User) *pb.User {
	return &pb.User{
		Id:          u.ID.String(),
		Email:       u.Email.Address,
		Name:        u.Name,
		Verified:    u.Verified,
		DateCreated: u.DateCreated.Format(time.RFC3339),
		DateUpdated: u.DateUpdated.Format(time.RFC3339),
	}
}
