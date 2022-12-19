package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"myiot/models"
	"myiot/user/rpc/internal/config"
	"myiot/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user.User
	AuthUser *user.UserAuthRequest
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
