package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	UserCookieDomin string
	Upstreams       map[string]string
}
