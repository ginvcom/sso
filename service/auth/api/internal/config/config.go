package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Etcd  discov.EtcdConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret         string // 生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire         int64  // jwt token有效期，单位：秒  1天=86400秒
		RememberAccessExpire int64  // jwt  记住登录时token有效期，单位：秒  3天=259200秒,  7天=604800秒
	}
	// NoAuthUrls是一个无需权限校验白名单
	// map的key是页面菜单的key, 数组元素是无需校验权限的请求(匹配请求方式和请求path)
	// 有条件的做成配置中心配置这个无需权限校验白名单
	NoAuthUrls []ServiceNoAuthData
}

type ServiceNoAuthData struct {
	SystemCode string // 服务的ServiceCode
	NoAuthData []NoAuthItem
}

type NoAuthItem struct {
	Menu string
	Urls []NoAuthUrl
}

type NoAuthUrl struct {
	Method string
	Path   string
}

type BasicContext string

const Token BasicContext = "token"
const UUID BasicContext = "uuid"
const Name BasicContext = "name"
const SystemCode BasicContext = "systemCode"
