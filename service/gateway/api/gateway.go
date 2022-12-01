package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sso/service/gateway/api/internal/config"
	"sso/service/gateway/api/internal/handler"
	"sso/service/gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	serviceNameField := logx.LogField{
		Key:   "serviceName",
		Value: c.Name,
	}
	envField := logx.LogField{
		Key:   "env",
		Value: c.Env,
	}
	logx.AddGlobalFields(serviceNameField, envField)

	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	var srv = &http.Server{
		Addr: address,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			referer := r.Header.Get("Referer")
			refererURL, err := url.Parse(referer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", refererURL.Scheme+"://"+refererURL.Host)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, x-client-system, x-client-service, x-client-env, x-client-uri")
			w.Header().Set("content-type", "application/json")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == http.MethodOptions {
				return
			}

			ctx, err := svc.NewServiceContext(c, r)
			if err != nil {
				if err == http.ErrNoCookie {
					w.WriteHeader(http.StatusUnauthorized)
					// w.Write([]byte(svc.ErrAccountLoginTimeout))
				} else {
					logx.Info(err)
					w.WriteHeader(http.StatusBadRequest)
					// w.Write([]byte(svc.ErrRequestParamsError))
				}
				return
			}

			if ctx.Meta.Action == svc.ActionImage {
				handler.ImageHandler(ctx, w, r)
			} else if ctx.Meta.Action == svc.ActionSignIn {
				handler.LoginHandler(ctx, w, r)
			} else if ctx.Meta.Action == svc.ActionSignOut {
				handler.SignOutHandler(ctx, w, r)
			} else {
				if ctx.Meta.Action != svc.ActionSessionMenus && ctx.Meta.Action != svc.ActionSessionMenuActions {
					//使用uri进行鉴权(白名单、单点操作)
					fmt.Println("ctx", ctx.Meta.URI)
					err = handler.VerifyHandler(ctx, w, r)
					// 鉴权通过, 附传用户uuid和用户名请求接口
					if err != nil {
						logx.Info(err)
						return
					}
				}

				body, err := svc.GetBody(ctx, r)
				if err != nil {
					logx.Info(err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				status, resp, err := ctx.Request(body)

				if err != nil {
					logx.Info(err)
					w.WriteHeader(http.StatusInternalServerError)
					// w.Write([]byte(err.Error()))
					return
				}

				if status != http.StatusOK {
					logx.Info(status)
					w.WriteHeader(status)
					// w.Write([]byte(http.StatusText(status)))
					return
				}

				w.Write(resp)
			}
		}),
	}
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigs
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("Server shutdown with error: ", err)
		}
		close(done)
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("Http server start failed", err)
	}
	<-done
}
