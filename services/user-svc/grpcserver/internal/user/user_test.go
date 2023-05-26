package user_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/adapter/local"
	"github.com/hausops/mono/services/user-svc/grpcserver/internal/user"
	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	svc := user.NewServer(local.NewUserRepository())

	email := gofakeit.Email()
	req := &pb.EmailRequest{Email: email}

	u, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Errorf("Create(%s) = error %v", email, err)
	}

	if _, err := uuid.Parse(u.Id); err != nil {
		t.Errorf("Id is not a uuid: %v", err)
	}

	if u.Email != email {
		t.Errorf("Email does not match. Got: %s; want: %s", u.Email, email)
	}

	if u.Verified {
		t.Error("User should be unverified")
	}

	if u.DateCreated == "" {
		t.Error("DateCreated is empty")
	}

	if u.DateUpdated != u.DateCreated {
		t.Error("DateUpdated should be the same as DateCreated at creation")
	}
}

func TestCreate_InvalidEmail(t *testing.T) {
	svc := user.NewServer(local.NewUserRepository())

	email := "invalid-email"
	req := &pb.EmailRequest{Email: email}

	_, err := svc.Create(context.Background(), req)
	if err == nil {
		t.Errorf("Create(%s) got no error", email)
	}

	s, _ := status.FromError(err)
	got := s.Code()
	want := codes.InvalidArgument
	if got != want {
		t.Errorf("Create(%s) got %s error code; want %s", email, got, want)
	}
}

func TestFindByEmail(t *testing.T) {
	svc := user.NewServer(local.NewUserRepository())

	email := gofakeit.Email()
	req := &pb.EmailRequest{Email: email}

	_, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create(%s) = error %v", email, err)
	}

	u, err := svc.FindByEmail(context.Background(), req)
	if err != nil {
		t.Errorf("FindByEmail(%s) = error %v", email, err)
	}

	if u.Email != email {
		t.Errorf("Email does not match. Got: %s; want: %s", u.Email, email)
	}
}

func TestFindByEmail_NotFound(t *testing.T) {
	svc := user.NewServer(local.NewUserRepository())

	// Setup: create user
	{
		email := gofakeit.Email()
		req := &pb.EmailRequest{Email: email}

		_, err := svc.Create(context.Background(), req)
		if err != nil {
			t.Fatalf("Create(%s) = error %v", email, err)
		}
	}

	email := "non-existing-email@hausops.com"
	req := &pb.EmailRequest{Email: email}

	_, err := svc.FindByEmail(context.Background(), req)
	if err == nil {
		t.Errorf("FindByEmail(%s) got no error", email)
	}

	s, _ := status.FromError(err)
	got := s.Code()
	want := codes.NotFound
	if got != want {
		t.Errorf("FindByEmail(%s) got %s error code; want %s", email, got, want)
	}
}
