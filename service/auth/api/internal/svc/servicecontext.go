package svc

import (
	"sso/service/auth/api/internal/config"
	"sso/service/auth/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       model.UserModel
	ObjectModel     model.ObjectModel
	UserToRoleModel model.UserToRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserModel:       model.NewUserModel(conn, c.CacheRedis),
		ObjectModel:     model.NewObjectModel(conn, c.CacheRedis),
		UserToRoleModel: model.NewUserToRoleModel(conn, c.CacheRedis),
	}
}
