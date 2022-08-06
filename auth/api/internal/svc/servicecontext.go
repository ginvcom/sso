package svc

import (
	"sso/auth/api/internal/config"
	"sso/auth/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	ObjectModel model.ObjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(conn),
		ObjectModel: model.NewObjectModel(conn),
	}
}
