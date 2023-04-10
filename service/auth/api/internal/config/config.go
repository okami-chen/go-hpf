package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	//Mysql struct {
	//	DataSource string
	//}
	//
	//CacheRedis cache.CacheConf
	JwtAuth struct {
		AccessSecret string
	}
}
