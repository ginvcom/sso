package main

import (
	"flag"
	"fmt"
	"net/http"

	"sso/service/auth/api/internal/config"
	"sso/service/auth/api/internal/handler"
	"sso/service/auth/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "X-Origin-Service")
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-uri")
	w.Header().Add("Access-Control-Allow-Headers", "X-Origin-Uri")
}
