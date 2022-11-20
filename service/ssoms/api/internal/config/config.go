package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Env     string
	NsqHost string
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string // 生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire int64  // jwt token有效期，单位：秒  1天=86400秒
	}
}

type BasicContext string

const UUID BasicContext = "uuid"
const Name BasicContext = "name"
