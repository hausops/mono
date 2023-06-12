package dapr

import (
	"context"

	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type userService struct {
	client pb.UserServiceClient
}

func NewUserService(conn grpc.ClientConnInterface) *userService {
	return &userService{client: pb.NewUserServiceClient(conn)}
}

// Ensure userService implements the pb.UserServiceClient interface.
var _ pb.UserServiceClient = (*userService)(nil)

func (s *userService) Create(ctx context.Context, in *pb.EmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
	return s.client.Create(
		s.withDaprContext(ctx),
		&pb.EmailRequest{Email: in.GetEmail()},
	)
}

func (s *userService) FindByEmail(ctx context.Context, in *pb.EmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
	return s.client.FindByEmail(
		s.withDaprContext(ctx),
		&pb.EmailRequest{Email: in.GetEmail()},
	)
}

func (s *userService) withDaprContext(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "user-svc")
}
