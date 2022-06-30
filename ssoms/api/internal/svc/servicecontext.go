package svc

import (
	"sso/ssoms/api/internal/config"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	RoleModel model.RoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
		RoleModel: model.NewRoleModel(conn),
	}
}
