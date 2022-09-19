package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/handler"
	"sso/service/ssoms/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ssoms-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "*"))
	defer server.Stop()

	// 全局中间件, 获取网关传递的用户基本信息
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			uuid := r.Header.Get("X-Origin-Uuid")
			name := r.Header.Get("X-Origin-Name")
			// 可能存在无需登录就能访问的页面
			if uuid == "" || name == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), config.UUID, uuid)
			ctx = context.WithValue(ctx, config.Name, name)
			next(w, r.WithContext(ctx))
		}
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "X-ginv-uri")
}
