// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	objectFieldNames          = builder.RawFieldNames(&Object{})
	objectRows                = strings.Join(objectFieldNames, ",")
)

type (
	objectModel interface {
		IsExist(ctx context.Context, args *ObjectIsExistArgs) (bool, error) // 用于新增、更新前的校验
		FindOneByUuid(ctx context.Context, uuid string) (*Object, error)
		FindOne(ctx context.Context, args *ObjectFindOneArgs) (*Object, error)
	}

	defaultObjectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ObjectListArgs struct {
		TopKey string
		ObjectName string
		Key string
		Status int64
		Typ int64
		ExcludeHide bool
	}

	ObjectIsExistArgs struct {
		TopKey string
		Key string
		Typ int64
		SubType int64
		ExcludeUUID string
	}

	Object struct {
		Id         int64     `db:"id"`
		Uuid       string    `db:"uuid"`
		ObjectName string    `db:"object_name"`
		Identifier     string    `db:"identifier"`
		Key        string    `db:"key"` // 操作对象的systemCode, 菜单的path, 操作的uri
		Sort       int64     `db:"sort"`
		Type       int64     `db:"type"`      // 类型: 1操作对象, 2菜单, 3操作(接口)
		SubType    int64     `db:"sub_type"`  // 子分类
		Extra      string    `db:"extra"`     // 扩展字段，建议封装成 JSON 字符串
		Icon       string    `db:"icon"`      // 图标
		Status     int64     `db:"status"`    // 状态
		Puuid      string    `db:"puuid"`     // 父级uuid
		TopKey     string    `db:"top_key"`   // 操作对象的所属systemCode
		IsDelete   int64     `db:"is_delete"` // 是否删除: 0正常, 1删除
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newObjectModel(conn sqlx.SqlConn) *defaultObjectModel {
	return &defaultObjectModel{
		conn:  conn,
		table: "`object`",
	}
}

func (m *defaultObjectModel) IsExist(ctx context.Context, args *ObjectIsExistArgs) (exist bool, err error) {
	where := "`top_key` = ? and `key` = ? and `type` = ? and `sub_type` = ? and `is_delete` = 0"
	placeholder :=[]interface{}{args.TopKey, args.Key, args.Typ, args.SubType}
	if args.ExcludeUUID != "" {
		where +=" and `uuid` != ?"
		placeholder = append(placeholder, args.ExcludeUUID)
	}
	query := fmt.Sprintf("select count(*) as count from %s where %s limit 1", m.table, where)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}
	var count int64
	err = stmt.QueryRowCtx(ctx, &count, placeholder...)

	if err !=nil{
		return
	}

	if count > 0 {
		exist = true
	}
	
	return
}

func (m *defaultObjectModel) FindOneByUuid(ctx context.Context, uuid string) (resp *Object,  err error) {
	query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", objectRows, m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	resp = new(Object)
	err = stmt.QueryRowCtx(ctx, resp, uuid)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}
	return
}

type ObjectFindOneArgs struct {
	Key string
	TopKey string
}

func (m *defaultObjectModel) FindOne(ctx context.Context, args *ObjectFindOneArgs) (resp *Object, err error) {
	where := "`status` = 1 and `is_delete` = 0"
	placeholder :=[]interface{}{args.TopKey}
	if args.TopKey != "" {
		where += " and `top_key`=?"
	} else {
		err = errors.New("topKey is required")
		return
	}
	if args.Key != "" {
		where += " and `key`=?"
		placeholder = append(placeholder, args.Key)
	}
	query := fmt.Sprintf("select %s from %s where %s limit 1", objectRows, m.table, where)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	resp = new(Object)
	err = stmt.QueryRowCtx(ctx, resp, placeholder...)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}

	return
}
