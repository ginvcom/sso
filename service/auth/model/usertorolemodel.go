package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserToRoleModel = (*customUserToRoleModel)(nil)

type (
	// UserToRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserToRoleModel.
	UserToRoleModel interface {
		userToRoleModel
		FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (roleUUIDArray *[]string, err error)
	}

	customUserToRoleModel struct {
		*defaultUserToRoleModel
	}
)

// NewUserToRoleModel returns a model for the database table.
func NewUserToRoleModel(conn sqlx.SqlConn, c cache.CacheConf) UserToRoleModel {
	return &customUserToRoleModel{
		defaultUserToRoleModel: newUserToRoleModel(conn, c),
	}
}

func (m *defaultUserToRoleModel) FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (roleUUIDArray *[]string, err error) {
	query, values, err := squirrel.Select("role_uuid").From(m.table).Where("is_delete=?", 0).Where("user_uuid=?", userUuid).ToSql()
	if err != nil {
		return
	}
	roleUUIDArray = new([]string)
	err = m.QueryRowsNoCacheCtx(ctx, roleUUIDArray, query, values...)

	return
}
