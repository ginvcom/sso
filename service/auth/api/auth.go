package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"sso/service/auth/api/internal/config"
	"sso/service/auth/api/internal/handler"
	"sso/service/auth/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "*"))
	defer server.Stop()

	// 添加全局日志字段
	serviceNameField := logx.LogField{
		Key:   "serviceName",
		Value: c.Name,
	}
	modeField := logx.LogField{
		Key:   "mode",
		Value: c.Mode,
	}
	logx.AddGlobalFields(serviceNameField, modeField)

	// 全局中间件, 获取网关传递的用户基本信息
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			uuid := r.Header.Get("x-origin-Uuid")
			name := url.QueryEscape(r.Header.Get("x-origin-name"))

			// 可能存在无需登录就能访问的页面
			// if uuid == "" || name == "" {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }
			uri := url.QueryEscape(r.Header.Get("x-origin-service"))
			systemCode := r.Header.Get("x-origin-system")
			token := r.Header.Get("x-origin-token")
			ctx := context.WithValue(r.Context(), config.UUID, uuid)
			ctx = context.WithValue(ctx, config.Name, name)
			ctx = context.WithValue(ctx, config.SystemCode, systemCode)
			ctx = context.WithValue(ctx, config.Token, token)
			uuidField := logx.LogField{
				Key:   "uuid",
				Value: uuid,
			}
			nameField := logx.LogField{
				Key:   "name",
				Value: name,
			}
			uriField := logx.LogField{
				Key:   "uri",
				Value: uri,
			}
			ctx = logx.WithFields(ctx, uriField, uuidField, nameField)
			next(w, r.WithContext(ctx))
		}
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	// 接受网关转发时的headers参数
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-service")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-uri")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-name")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-system")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-token")
}
