package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	Env string
	rest.RestConf
	UserCookieDomin string
	Upstreams       map[string]string
}
