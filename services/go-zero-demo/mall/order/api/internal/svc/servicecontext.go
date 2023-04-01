package svc

import (
	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
