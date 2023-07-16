package mock

import (
	"context"
	"errors"

	"github.com/hausops/mono/services/user-svc/adapter/mock"
	"github.com/hausops/mono/services/user-svc/grpcserver"
	"github.com/hausops/mono/services/user-svc/pb"
	"google.golang.org/grpc"
)

// userServiceClient implements pb.UserServiceClient as a mock so services
// depending on user-svc can use it as a test client in tests.
type userServiceClient struct {
	server pb.UserServiceServer
}

// NewUserServiceClient creates a userServiceClient using mock user.Repository.
func NewUserServiceClient() *userServiceClient {
	repo := mock.NewUserRepository()
	server := grpcserver.NewUserServer(repo)
	return &userServiceClient{server: server}
}

var _ pb.UserServiceClient = (*userServiceClient)(nil)

func (c *userServiceClient) Create(ctx context.Context, in *pb.EmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
	if len(opts) > 0 {
		return nil, errors.New("mock.UserServiceClient does not support grpc.CallOption")
	}
	return c.server.Create(ctx, in)
}

func (c *userServiceClient) FindByEmail(ctx context.Context, in *pb.EmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
	if len(opts) > 0 {
		return nil, errors.New("mock.UserServiceClient does not support grpc.CallOption")
	}
	return c.server.FindByEmail(ctx, in)
}
