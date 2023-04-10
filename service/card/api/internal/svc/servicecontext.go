package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"hpf/common/middleware"
	"hpf/service/card/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
