package main

import (
	"flag"
	"fmt"
	"net/http"

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

	serviceNameField := logx.LogField{
		Key:   "serviceName",
		Value: c.Name,
	}
	envField := logx.LogField{
		Key:   "env",
		Value: c.Env,
	}
	logx.AddGlobalFields(serviceNameField, envField)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-service")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-uri")
	// w.Header().Add("Access-Control-Allow-Headers", "x-origin-Uri")
}
