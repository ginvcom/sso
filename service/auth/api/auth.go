package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"sso/service/auth/api/internal/config"
	"sso/service/auth/api/internal/handler"
	"sso/service/auth/api/internal/svc"

	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth-api.yaml", "the config file")

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

	// server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "*"))
	defer server.Stop()

	writer := logx.NewWriter(NewNSQWriter(producer))
	logx.SetWriter(writer)

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
