package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserToRoleModel = (*customUserToRoleModel)(nil)

type (
	// UserToRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserToRoleModel.
	UserToRoleModel interface {
		userToRoleModel
	}

	customUserToRoleModel struct {
		*defaultUserToRoleModel
	}
)

// NewUserToRoleModel returns a model for the database table.
func NewUserToRoleModel(conn sqlx.SqlConn) UserToRoleModel {
	return &customUserToRoleModel{
		defaultUserToRoleModel: newUserToRoleModel(conn),
	}
}
