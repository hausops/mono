package grpcserver_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hausops/mono/services/user-svc/adapter/mock"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/hausops/mono/services/user-svc/grpcserver"
	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())

	email := gofakeit.Email()
	req := &pb.EmailRequest{Email: email}

	u, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Errorf("Create(%s) error = %v", email, err)
	}

	if _, err := user.ParseID(u.Id); err != nil {
		t.Errorf("Id is invalid: %v", err)
	}

	if u.Email != email {
		t.Errorf("Email does not match; got %s, want %s", u.Email, email)
	}

	if u.DateCreated == "" {
		t.Error("DateCreated is empty")
	}

	if u.DateUpdated != u.DateCreated {
		t.Error("DateUpdated should be the same as DateCreated at creation")
	}
}

func TestCreate_InvalidEmail(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())

	for _, email := range []string{"", "invalid-email"} {
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
}

func TestCreate_EmailTaken(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())
	u := mustCreateUser(t, svc)

	req := &pb.EmailRequest{Email: u.Email}

	_, err := svc.Create(context.Background(), req)
	if err == nil {
		t.Errorf("Create(%s) got no error", u.Email)
	}

	s, _ := status.FromError(err)
	got := s.Code()
	want := codes.AlreadyExists
	if got != want {
		t.Errorf("Create(%s) got %s error code; want %s", u.Email, got, want)
	}
}

func TestFindByEmail(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())
	u := mustCreateUser(t, svc)

	req := &pb.EmailRequest{Email: u.Email}

	found, err := svc.FindByEmail(context.Background(), req)
	if err != nil {
		t.Errorf("FindByEmail(%s) error = %v", u.Email, err)
	}

	got := found.Email
	want := u.Email
	if got != want {
		t.Errorf("Email does not match; got %s, want %s", got, want)
	}
}

func TestFindByEmail_InvalidEmail(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())
	mustCreateUser(t, svc)

	email := "invalid-email"
	req := &pb.EmailRequest{Email: email}

	_, err := svc.FindByEmail(context.Background(), req)

	s, _ := status.FromError(err)
	got := s.Code()
	want := codes.InvalidArgument
	if got != want {
		t.Errorf("FindByEmail(%s) got %s error code; want %s", email, got, want)
	}
}

func TestFindByEmail_NotFound(t *testing.T) {
	svc := grpcserver.NewUserServer(mock.NewUserRepository())
	mustCreateUser(t, svc)

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

func mustCreateUser(t *testing.T, svc pb.UserServiceServer) *pb.User {
	t.Helper()
	email := gofakeit.Email()
	req := &pb.EmailRequest{Email: email}

	u, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create(%s) error = %v", email, err)
	}
	return u
}
