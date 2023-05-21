package user

import (
	"github.com/hausops/mono/services/user-svc/domain/user"
	"github.com/hausops/mono/services/user-svc/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	svc *user.Service
}

func NewServer(svc *user.Service) *server {
	return &server{svc: svc}
}
