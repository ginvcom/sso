// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userInsertFields   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`is_delete`"), ",")
	userUpdateFields = strings.Join(stringx.Remove(userFieldNames, "`id`", "uuid", "`create_time`", "`is_delete`"), "=?,") + "=?"
	userUpdateFieldsWithoutPassword = strings.Join(stringx.Remove(userFieldNames, "`id`", "uuid", "`create_time`", "`password`", "`is_delete`"), "=?,") + "=?"
)

type (
	userModel interface {
		ListCount(ctx context.Context, args *UserListArgs) (int64, error)
		ListData(ctx context.Context, args *UserListArgs) (*[]User, error)
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOneByUuid(ctx context.Context, uuid string) (*User, error)
		Update(ctx context.Context, newData *User) error
		Delete(ctx context.Context, uuid string) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserListArgs struct {
		Name string
		Mobile string
		Page int64
		PageSize int64
	}

	User struct {
		Id         int64     `db:"id"`
		Uuid       string    `db:"uuid"`
		Avatar     string    `db:"avatar"`    // 头像
		Name       string    `db:"name"`      // 姓名
		Mobile     string    `db:"mobile"`    // 手机号
		Password   string    `db:"password"`  // 密码
		Gender     int64     `db:"gender"`    // 性别: 1男, 2女, 3未知
		Birth      time.Time `db:"birth"`     // 生日
		Status     int64     `db:"status"`    // 状态: 1正常
		IsDelete   int64     `db:"is_delete"` // 是否删除: 0正常, 1删除
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}

)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (args *UserListArgs) getListConditions () (where string, placeholder []interface{}) {
	where = "`is_delete` = 0"
	if args.Mobile != "" {
		where +=" and mobile = ?"
		placeholder = append(placeholder, args.Mobile)
	}
	if args.Name != "" {
		where +=" and name like ?"
		placeholder = append(placeholder, args.Name + "%")
	}
	return
}

func (m *defaultUserModel) ListCount(ctx context.Context, args *UserListArgs) (count int64, err error) {
	var placeholder []interface{}
	where, placeholder := args.getListConditions()
	query := fmt.Sprintf("select count(*) as count from %s where %s limit 1", m.table, where)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}

	err = stmt.QueryRowCtx(ctx, &count, placeholder...)
	
	return
}

func (m *defaultUserModel) ListData(ctx context.Context, args *UserListArgs) (resp *[]User, err error) {
	var placeholder []interface{}
	where, placeholder := args.getListConditions()
	offset:= (args.Page - 1) * args.PageSize
	placeholder = append(placeholder, args.PageSize, offset)
	query := fmt.Sprintf("select %s from %s where %s limit ? offset ?", userRows, m.table, where)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	resp = new([]User)
	err = stmt.QueryRowsCtx(ctx, resp, placeholder...)

	return resp, err
}

func (m *defaultUserModel) Delete(ctx context.Context, uuid string) (err error) {
	updateTime := time.Now().Local()
	query := fmt.Sprintf("update %s set `is_delete`=1, `update_time`= ? where `uuid` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}

	res, err := stmt.ExecCtx(ctx, updateTime, uuid)
	if err!=nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err !=nil{
		return
	}

	if rowsAffected == 0 {
		err = errors.New("no rows affected")
	}

	return err
}

func (m *defaultUserModel) FindOneByUuid(ctx context.Context, uuid string) (u *User, err error) {
	query := fmt.Sprintf("select %s from %s where `is_delete` = 0 and `uuid` = ? limit 1", userRows, m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	u = new(User)
	err = stmt.QueryRowCtx(ctx, u, uuid)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}

	return
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (res sql.Result, err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInsertFields)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err !=nil{
		return
	}

	res, err = stmt.ExecCtx(ctx, data.Uuid, data.Avatar, data.Name, data.Mobile, data.Password, data.Gender, data.Birth, data.Status)

	return
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) (err error) {
	updateTime := time.Now().Local()
	var placeholder []interface{}
	var stmt sqlx.StmtSession
	if newData.Password == "" {
		query := fmt.Sprintf("update %s set %s where `is_delete` = 0 and `uuid` = ?", m.table, userUpdateFieldsWithoutPassword)
		stmt, err = m.conn.PrepareCtx(ctx, query)
		
	} else {
		query := fmt.Sprintf("update %s set %s where `is_delete` = 0 and `uuid` = ?", m.table, userUpdateFields)
		stmt, err = m.conn.PrepareCtx(ctx, query)
		placeholder = append(placeholder, newData.Uuid, newData.Avatar, newData.Name, newData.Mobile, newData.Password, newData.Gender, newData.Birth, newData.Status, updateTime, newData.Uuid)
	}
	if err !=nil{
		return
	}

	res, err := stmt.ExecCtx(ctx, placeholder...)
	if err !=nil{
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err !=nil{
		return
	}

	if rowsAffected == 0 {
		err = errors.New("no rows affected")
	}

	return err
}
