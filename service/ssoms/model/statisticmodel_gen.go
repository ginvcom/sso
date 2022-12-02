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
	statisticFieldNames          = builder.RawFieldNames(&Statistic{})
	statisticRows                = strings.Join(statisticFieldNames, ",")
	statisticRowsExpectAutoSet   = strings.Join(stringx.Remove(statisticFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	statisticRowsWithPlaceHolder = strings.Join(stringx.Remove(statisticFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	statisticModel interface {
		Insert(ctx context.Context, data *Statistic) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Statistic, error)
		FindOneByMonth(ctx context.Context, month string) (*Statistic, error)
		Update(ctx context.Context, data *Statistic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStatisticModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Statistic struct {
		Id               int64     `db:"id"`
		Month            string    `db:"month"`             // 月份, YYYY-MM格式
		UserAmount       int64     `db:"user_amount"`       // 用户数量
		RoleAmount       int64     `db:"role_amount"`       // 角色数量
		SystemAmount     int64     `db:"system_amount"`     // 系统数量
		MenuAmount       int64     `db:"menu_amount"`       // 菜单数量
		ActionAmount     int64     `db:"action_amount"`     // 操作数量
		PermissionAmount int64     `db:"permission_amount"` // 授权数量
		CreateTime       time.Time `db:"create_time"`
		UpdateTime       time.Time `db:"update_time"`
	}
)

func newStatisticModel(conn sqlx.SqlConn) *defaultStatisticModel {
	return &defaultStatisticModel{
		conn:  conn,
		table: "`statistic`",
	}
}

func (m *defaultStatisticModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultStatisticModel) FindOne(ctx context.Context, id int64) (*Statistic, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", statisticRows, m.table)
	var resp Statistic
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

func (m *defaultStatisticModel) FindOneByMonth(ctx context.Context, month string) (*Statistic, error) {
	var resp Statistic
	query := fmt.Sprintf("select %s from %s where `month` = ? limit 1", statisticRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, month)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStatisticModel) Insert(ctx context.Context, data *Statistic) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, statisticRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Month, data.UserAmount, data.RoleAmount, data.SystemAmount, data.MenuAmount, data.ActionAmount, data.PermissionAmount)
	return ret, err
}

func (m *defaultStatisticModel) Update(ctx context.Context, newData *Statistic) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, statisticRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Month, newData.UserAmount, newData.RoleAmount, newData.SystemAmount, newData.MenuAmount, newData.ActionAmount, newData.PermissionAmount, newData.Id)
	return err
}

func (m *defaultStatisticModel) tableName() string {
	return m.table
}
