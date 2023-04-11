package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"hpf/common/middleware"
	"hpf/pkg/starter"
	"hpf/service/card/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := starter.NewMysqlConn(c.Mysql)
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		DbEngin: conn,
	}
}
