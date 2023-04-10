package config

import (
	"github.com/zeromicro/go-zero/rest"
	"hpf/common/config"
)

type Config struct {
	rest.RestConf
	//Mysql struct {
	//	DataSource string
	//}
	//
	//CacheRedis cache.CacheConf
	JwtAuth config.JwtAuth
}
