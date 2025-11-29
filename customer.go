package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Nozomi9967/wmss-customer-api/internal/config"
	"github.com/Nozomi9967/wmss-customer-api/internal/handler"
	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/middleware"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/customer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//middleware.InitSimpleLogger()

	// 关闭统计日志（去掉 p2c 那些 stat 日志）
	logx.DisableStat()

	server := rest.MustNewServer(c.Rest,
		rest.WithNotAllowedHandler(http.NotFoundHandler()),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	jwtMiddleware := middleware.NewJwtAuthMiddleware(ctx)
	ctx.JwtAuthMiddleware = jwtMiddleware.Handle

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Rest.Host, c.Rest.Port)
	server.Start()
}
