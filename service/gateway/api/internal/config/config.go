package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	Env     string
	NsqHost string
	rest.RestConf
	AllowOrigin     string
	UserCookieDomin string
	Upstreams       map[string]string
}
