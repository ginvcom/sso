package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		RowBuilder() squirrel.SelectBuilder
		CustomFindOne(ctx context.Context, selectBuilder squirrel.SelectBuilder, mobile string) (u *User, err error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

// export logic
func (m *defaultUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRows).From(m.table)
}

func (m *defaultUserModel) CustomFindOne(ctx context.Context, selectBuilder squirrel.SelectBuilder, mobile string) (u *User, err error) {
	query, values, err := selectBuilder.Where("mobile=?", mobile).Where("`status`=?", 1).Where("is_delete=?", 0).Limit(1).ToSql()
	if err != nil {
		return
	}
	u = new(User)
	err = m.QueryRowNoCacheCtx(ctx, u, query, values...)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}

	return
}

func (m *defaultUserModel) CustomFindOneByUuid(ctx context.Context, selectBuilder squirrel.SelectBuilder, uuid string) (u *User, err error) {
	query, values, err := selectBuilder.Where("uuid=?", uuid).Where("is_delete=?", 0).Limit(1).ToSql()
	if err != nil {
		return
	}
	u = new(User)
	err = m.QueryRowNoCacheCtx(ctx, u, query, values...)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}

	return
}
