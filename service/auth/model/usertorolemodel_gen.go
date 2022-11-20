// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userToRoleFieldNames          = builder.RawFieldNames(&UserToRole{})
	userToRoleRows                = strings.Join(userToRoleFieldNames, ",")
	userToRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(userToRoleFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	userToRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(userToRoleFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheUserToRoleIdPrefix               = "cache:userToRole:id:"
	cacheUserToRoleUserUuidRoleUuidPrefix = "cache:userToRole:userUuid:roleUuid:"
)

type (
	userToRoleModel interface {
		Insert(ctx context.Context, data *UserToRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserToRole, error)
		FindOneByUserUuidRoleUuid(ctx context.Context, userUuid string, roleUuid string) (*UserToRole, error)
		Update(ctx context.Context, data *UserToRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserToRoleModel struct {
		sqlc.CachedConn
		table string
	}

	UserToRole struct {
		Id         int64     `db:"id"`
		UserUuid   string    `db:"user_uuid"`
		RoleUuid   string    `db:"role_uuid"`
		IsDelete   int64     `db:"is_delete"` // 是否删除: 0正常, 1删除
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserToRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserToRoleModel {
	return &defaultUserToRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_to_role`",
	}
}

func (m *defaultUserToRoleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userToRoleIdKey := fmt.Sprintf("%s%v", cacheUserToRoleIdPrefix, id)
	userToRoleUserUuidRoleUuidKey := fmt.Sprintf("%s%v:%v", cacheUserToRoleUserUuidRoleUuidPrefix, data.UserUuid, data.RoleUuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userToRoleIdKey, userToRoleUserUuidRoleUuidKey)
	return err
}

func (m *defaultUserToRoleModel) FindOne(ctx context.Context, id int64) (*UserToRole, error) {
	userToRoleIdKey := fmt.Sprintf("%s%v", cacheUserToRoleIdPrefix, id)
	var resp UserToRole
	err := m.QueryRowCtx(ctx, &resp, userToRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userToRoleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserToRoleModel) FindOneByUserUuidRoleUuid(ctx context.Context, userUuid string, roleUuid string) (*UserToRole, error) {
	userToRoleUserUuidRoleUuidKey := fmt.Sprintf("%s%v:%v", cacheUserToRoleUserUuidRoleUuidPrefix, userUuid, roleUuid)
	var resp UserToRole
	err := m.QueryRowIndexCtx(ctx, &resp, userToRoleUserUuidRoleUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_uuid` = ? and `role_uuid` = ? limit 1", userToRoleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userUuid, roleUuid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserToRoleModel) Insert(ctx context.Context, data *UserToRole) (sql.Result, error) {
	userToRoleIdKey := fmt.Sprintf("%s%v", cacheUserToRoleIdPrefix, data.Id)
	userToRoleUserUuidRoleUuidKey := fmt.Sprintf("%s%v:%v", cacheUserToRoleUserUuidRoleUuidPrefix, data.UserUuid, data.RoleUuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userToRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserUuid, data.RoleUuid, data.IsDelete)
	}, userToRoleIdKey, userToRoleUserUuidRoleUuidKey)
	return ret, err
}

func (m *defaultUserToRoleModel) Update(ctx context.Context, newData *UserToRole) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userToRoleIdKey := fmt.Sprintf("%s%v", cacheUserToRoleIdPrefix, data.Id)
	userToRoleUserUuidRoleUuidKey := fmt.Sprintf("%s%v:%v", cacheUserToRoleUserUuidRoleUuidPrefix, data.UserUuid, data.RoleUuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userToRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserUuid, newData.RoleUuid, newData.IsDelete, newData.Id)
	}, userToRoleIdKey, userToRoleUserUuidRoleUuidKey)
	return err
}

func (m *defaultUserToRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserToRoleIdPrefix, primary)
}

func (m *defaultUserToRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userToRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserToRoleModel) tableName() string {
	return m.table
}
