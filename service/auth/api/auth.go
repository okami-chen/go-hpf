package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hpf/common/errorx"
	"hpf/service/auth/api/internal/config"
	"hpf/service/auth/api/internal/handler"
	"hpf/service/auth/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CustomCodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	stat.DisableLog()
	load.DisableLog()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
