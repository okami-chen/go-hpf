package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hpf/common/errorx"
	"net/http"

	"hpf/service/card/api/internal/config"
	"hpf/service/card/api/internal/handler"
	"hpf/service/card/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/card.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CustomCodeError:
			return http.StatusOK, e.Data(ctx)
		default:
			return http.StatusInternalServerError, nil
		}
	})
	stat.DisableLog()
	load.DisableLog()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
