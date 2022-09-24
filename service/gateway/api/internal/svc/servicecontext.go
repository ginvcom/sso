package svc

import (
	"net/http"
	"sso/service/gateway/api/internal/config"
)

const (
	ErrServiceNotExits        = "servcie not exits"                    // 服务不存在
	ErrParsingHtml            = "error parsing html"                   // html 解析出错
	ErrRequestParamsError     = "unrecognized request parameters"      // 无法识别请求参数
	ErrAccountLoginTimeout    = "account login timeout"                // 账号登录超时
	ErrServiceNotResponding   = "service not responding"               // 服务未响应
	ErrServiceUnknownResponse = "service returned an unknown response" // 服务返回未知响应
)

type ServiceContext struct {
	Config config.Config
	Meta   Meta
}

func NewServiceContext(c config.Config, r *http.Request) (*ServiceContext, error) {
	meta, err := GetMeta(r)
	if err != nil {
		return nil, err
	}

	return &ServiceContext{
		Config: c,
		Meta:   meta,
	}, nil
}
