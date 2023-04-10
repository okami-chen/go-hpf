package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"hpf/common/middleware"
	"hpf/pkg/db"
	"hpf/service/card/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := db.NewMysqlConn(c.Mysql)
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		DbEngin: conn,
	}
}
