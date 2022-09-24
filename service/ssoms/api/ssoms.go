package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/handler"
	"sso/service/ssoms/api/internal/svc"

	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ssoms-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "*"))
	defer server.Stop()

	writer := logx.NewWriter(NewNSQWriter(producer))
	logx.SetWriter(writer)

	// 全局中间件, 获取网关传递的用户基本信息
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			uuid := r.Header.Get("x-origin-Uuid")
			name := r.Header.Get("x-origin-name")
			// 可能存在无需登录就能访问的页面
			if uuid == "" || name == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			uri := r.Header.Get("x-origin-service")
			ctx := context.WithValue(r.Context(), config.UUID, uuid)
			ctx = context.WithValue(ctx, config.Name, name)
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
	w.Header().Add("Access-Control-Allow-Headers", "x-origin-service")
}

type NSQWriter struct {
	Producer *nsq.Producer
}

func NewNSQWriter(producer *nsq.Producer) *NSQWriter {
	return &NSQWriter{
		Producer: producer,
	}
}

func (w *NSQWriter) Write(p []byte) (n int, err error) {
	logTopic := "go-zero-log"
	// 日志写入有换行符, 使用trim去掉.
	messageBody := strings.TrimSpace(string(p))
	if err := w.Producer.Publish(logTopic, []byte(messageBody)); err != nil {
		return 0, err
	}

	return len(p), nil
}
