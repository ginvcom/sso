package svc

import (
	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       model.UserModel
	RoleModel       model.RoleModel
	UserToRoleModel model.UserToRoleModel
	ObjectModel     model.ObjectModel
	ObjectLogModel  model.ObjectLogModel
	PermissionModel model.PermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserModel:       model.NewUserModel(conn),
		RoleModel:       model.NewRoleModel(conn),
		UserToRoleModel: model.NewUserToRoleModel(conn),
		ObjectModel:     model.NewObjectModel(conn),
		ObjectLogModel:  model.NewObjectLogModel(conn),
		PermissionModel: model.NewPermissionModel(conn),
	}
}
