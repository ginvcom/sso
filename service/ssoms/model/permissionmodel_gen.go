// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	permissionFieldNames          = builder.RawFieldNames(&Permission{})
	permissionRows                = strings.Join(permissionFieldNames, ",")
	permissionRowsExpectAutoSet   = strings.Join(stringx.Remove(permissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	permissionRowsWithPlaceHolder = strings.Join(stringx.Remove(permissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
	permissionInsertFeilds   = strings.Join(stringx.Remove(permissionFieldNames, "`id`", "`create_time`", "`is_delete`"), ",")
)

type (
	permissionModel interface {
		Insert(ctx context.Context, data *Permission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Permission, error)
		Update(ctx context.Context, data *Permission) error
		Delete(ctx context.Context, id int64) error
		PermissionByRoleUUID(ctx context.Context, roleUuid, topKey string) (*[]Permission, error)
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		TransDeletePermission(ctx context.Context, session sqlx.Session, roleUUID, topKey string, updateTime time.Time) error
		TransAddPermission(ctx context.Context, session sqlx.Session, dataArray []Permission, updateTime time.Time) error
		CountPermisson(ctx context.Context) (int64, error)
	}

	defaultPermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Permission struct {
		Id         int64     `db:"id"`
		RoleUuid   string    `db:"role_uuid"`
		ObjectUuid string    `db:"object_uuid"`
		Type       int64     `db:"type"`
		TopKey     string    `db:"top_key"`
		IsDelete   int64     `db:"is_delete"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newPermissionModel(conn sqlx.SqlConn) *defaultPermissionModel {
	return &defaultPermissionModel{
		conn:  conn,
		table: "`permission`",
	}
}

func (m *defaultPermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPermissionModel) FindOne(ctx context.Context, id int64) (*Permission, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", permissionRows, m.table)
	var resp Permission
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPermissionModel) Insert(ctx context.Context, data *Permission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, permissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleUuid, data.ObjectUuid, data.Type, data.TopKey, data.IsDelete)
	return ret, err
}

func (m *defaultPermissionModel) Update(ctx context.Context, data *Permission) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, permissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.RoleUuid, data.ObjectUuid, data.Type, data.TopKey, data.IsDelete, data.Id)
	return err
}

func (m *defaultPermissionModel) PermissionByRoleUUID(ctx context.Context, roleUuid, topKey string) (resp *[]Permission, err error) {
	query := fmt.Sprintf("select %s from %s where `is_delete` = 0 and `role_uuid` = ? and `top_key` = ?", permissionRows, m.table)
	fmt.Println(query)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	resp = new([]Permission)
	err = stmt.QueryRowsCtx(ctx, resp, roleUuid, topKey)

	return resp, err
}

// 提供给logic开启事务用
func (m *defaultPermissionModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func (ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultPermissionModel) TransDeletePermission(ctx context.Context, session sqlx.Session, roleUUID, topKey string, updateTime time.Time) (err error) {
	query := fmt.Sprintf("update %s set `is_delete`=1, `update_time`= ? where `role_uuid` = ? and `top_key` = ?", m.table)
	stmt, err:= session.PrepareCtx(ctx, query)
	if err != nil{
		return
	}
	_, err = stmt.ExecCtx(ctx, updateTime, roleUUID, topKey)
	return
}

func (m *defaultPermissionModel) TransAddPermission(ctx context.Context, session sqlx.Session, dataArray []Permission, updateTime time.Time) (err error) {
	fields := ""
	var placeholder []interface{}
	for i, item:= range dataArray {
		if i > 0 {
			fields += ", "
		}
		fields += "(?, ?, ?, ?, ?)"
 		placeholder = append(placeholder, item.RoleUuid, item.ObjectUuid, item.Type, item.TopKey, updateTime)
	}
	
	query := fmt.Sprintf("insert into %s(%s) values %s on duplicate key update `is_delete`= 0", m.table, permissionInsertFeilds, fields)
	stmt, err:= session.PrepareCtx(ctx, query)
	if err != nil{
		return
	}
	_, err = stmt.ExecCtx(ctx, placeholder...)
	return
}

// 首页授权数量统计
func (m *defaultPermissionModel) CountPermisson(ctx context.Context) (count int64, err error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_delete`= 0", m.table)
	err = m.conn.QueryRowCtx(ctx, &count, query)
	return
}

// func (m *defaultPermissionModel) tableName() string {
// 	return m.table
// }
