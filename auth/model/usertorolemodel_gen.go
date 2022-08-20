// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	userToRoleModel interface {
		FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (*[]string, error)
	}

	defaultUserToRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserToRole struct {
		Id         int64     `db:"id"`
		UserUuid   string     `db:"user_uuid"`
		RoleUuid   string     `db:"role_uuid"`
		IsDelete   int64     `db:"is_delete"` // 是否删除: 0正常, 1删除
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserToRoleModel(conn sqlx.SqlConn) *defaultUserToRoleModel {
	return &defaultUserToRoleModel{
		conn:  conn,
		table: "`user_to_role`",
	}
}

func (m *defaultUserToRoleModel) FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (roleUUIDArray *[]string, err error) {
	query := fmt.Sprintf("select `role_uuid` from %s where `is_delete` = 0 and `user_uuid` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	roleUUIDArray = new([]string)
	err = stmt.QueryRowsCtx(ctx, roleUUIDArray, userUuid)

	return
}
