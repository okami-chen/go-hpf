package config

import (
	"github.com/zeromicro/go-zero/rest"
	"hpf/common/config"
)

type Config struct {
	rest.RestConf
	JwtAuth config.JwtAuth
	Mysql   config.Mysql
}
