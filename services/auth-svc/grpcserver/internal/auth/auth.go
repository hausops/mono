package auth

import (
	"context"
	"fmt"
	"net/mail"
	"text/template"

	"github.com/hausops/mono/services/auth-svc/domain/credential"
	"github.com/hausops/mono/services/auth-svc/domain/email"
	"github.com/hausops/mono/services/auth-svc/domain/verification"
	"github.com/hausops/mono/services/auth-svc/pb"
	userpb "github.com/hausops/mono/services/user-svc/pb"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedAuthServer
	logger            *zap.Logger
	userSvc           userpb.UserServiceClient
	credentialRepo    credential.Repository
	verificationRepo  verification.Repository
	verificationEmail *verificationEmailSender
}

func NewServer(
	logger *zap.Logger,
	userSvc userpb.UserServiceClient,
	credentialRepo credential.Repository,
	verificationRepo verification.Repository,
	email email.Dispatcher,
) *server {
	return &server{
		logger:            logger,
		userSvc:           userSvc,
		credentialRepo:    credentialRepo,
		verificationRepo:  verificationRepo,
		verificationEmail: newVerificationEmailSender(email),
	}
}

func (s *server) SignUp(ctx context.Context, r *pb.SignUpRequest) (*emptypb.Empty, error) {
	email, err := mail.ParseAddress(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	// TODO: password policy + strength
	if r.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}

	// TODO: rollback the user creation if any steps below fail
	_, err = s.userSvc.Create(ctx, &userpb.EmailRequest{Email: email.Address})
	if err != nil {
		switch st, _ := status.FromError(err); st.Code() {
		case codes.AlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, fmt.Errorf("userSvc.Create(%s): %w", email.Address, err)
		}
	}

	hashedPassword, err := hashPassword(r.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	err = s.credentialRepo.Upsert(
		ctx,
		credential.Credential{
			Email:    *email,
			Password: hashedPassword,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("save credential %s: %w", email.Address, err)
	}

	verificationToken := verification.GenerateToken()
	err = s.verificationEmail.Send(ctx, *email, verificationToken)
	if err != nil {
		return nil, fmt.Errorf("send verification email to %s: %w", email.Address, err)
	}

	err = s.verificationRepo.Upsert(
		ctx,
		verification.PendingVerification{
			Email: *email,
			Token: verificationToken,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("save verification %s: %w", email.Address, err)
	}

	return new(emptypb.Empty), nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

type verificationEmailSender struct {
	email    email.Dispatcher
	template *template.Template
}

func newVerificationEmailSender(email email.Dispatcher) *verificationEmailSender {
	const templateFilename = "grpcserver/internal/auth/verification-email.txt"
	tmpl := template.Must(template.ParseFiles(templateFilename))
	return &verificationEmailSender{
		email:    email,
		template: tmpl,
	}
}

func (s *verificationEmailSender) Send(ctx context.Context, to mail.Address, token verification.Token) error {
	subject := "Verify your email to start using HausOps"
	body := fmt.Sprintf("Verify your email address: https://auth.hausops.com/verify-email?t=%s", token)
	return s.email.Send(ctx, to, subject, body)
}
