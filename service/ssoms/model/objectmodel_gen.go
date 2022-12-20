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
	objectFieldNames          = builder.RawFieldNames(&Object{})
	objectRows                = strings.Join(objectFieldNames, ",")
	objectInsertFields   = strings.Join(stringx.Remove(objectFieldNames, "`id`", "`create_time`", "`update_time`", "`is_delete`"), ",")
	objectUpdateFields = strings.Join(stringx.Remove(objectFieldNames, "`id`", "`uuid`",  "`top_key`",  "`create_time`", "`is_delete`"), "=?,") + "=?"
)

type (
	objectModel interface {
		ListData(ctx context.Context, args *ObjectListArgs) (*[]Object, error)
		Insert(ctx context.Context, data *Object) (sql.Result, error)
		IsExist(ctx context.Context, args *ObjectIsExistArgs) (bool, error) // 用于新增、更新前的校验
		FindOneByUuid(ctx context.Context, uuid string) (*Object, error)
		Update(ctx context.Context, data *Object) error
		Delete(ctx context.Context, uuid string) error
		CountByType(ctx context.Context) (*[]ObjectTypeCountItem, error)
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

	ObjectTypeCountItem struct{
		Typ    int64    `db:"type"`      // 类型: 1操作对象, 2菜单, 3操作(接口)
		Count   int64    `db:"count"`    // 数量
	}
)

func newObjectModel(conn sqlx.SqlConn) *defaultObjectModel {
	return &defaultObjectModel{
		conn:  conn,
		table: "`object`",
	}
}

func (args *ObjectListArgs) getListConditions () (where string, placeholder []interface{}) {
	where = "`is_delete` = 0"
	if args.Status != 0 {
		where +=" and `status` = ?"
		if (args.Status == -1) {
			args.Status = 0
		}
		placeholder = append(placeholder, args.Status)
	}

	if args.Typ != 0 {
		where +=" and `type` = ?"
		placeholder = append(placeholder, args.Typ)
		if args.TopKey != "" {
			where +=" and top_key = ?"
			placeholder = append(placeholder, args.TopKey)
		}
		if args.ExcludeHide {
			where +=" and sub_type !=3"
		}
	} else {
		if args.TopKey != "" {
			where +=" and top_key = ? and type != 1"
			placeholder = append(placeholder, args.TopKey)
		} else {
			where +=" and `type` = 1"
		}
	}
	if args.ObjectName != "" {
		where +=" and object_name = ?"
		placeholder = append(placeholder, args.ObjectName)
	}
	if args.Key != "" {
		where +=" and `key` like ?"
		placeholder = append(placeholder, args.Key + "%")
	}
	return
}

func (m *defaultObjectModel) ListData(ctx context.Context, args *ObjectListArgs) (resp *[]Object, err error) {
	var placeholder []interface{}
	where, placeholder := args.getListConditions()
	query := fmt.Sprintf("select %s from %s where %s order by sort, create_time", objectRows, m.table, where)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	resp = new([]Object)
	err = stmt.QueryRowsCtx(ctx, resp, placeholder...)

	return resp, err
}

func (m *defaultObjectModel) Delete(ctx context.Context, uuid string) (err error) {
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
	return

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

func (m *defaultObjectModel) FindOneByUuid(ctx context.Context, uuid string) (resp *Object, err error) {
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

func (m *defaultObjectModel) Insert(ctx context.Context, data *Object) (res sql.Result, err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, objectInsertFields)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}
	res, err= stmt.ExecCtx(
		ctx,
		data.Uuid,
		data.ObjectName,
		data.Identifier,
		data.Key, 
		data.Sort,
		data.Type,
		data.SubType,
		data.Extra,
		data.Icon,
		data.Status,
		data.Puuid,
		data.TopKey,
	)
	
	return
}

func (m *defaultObjectModel) Update(ctx context.Context, newData *Object) (err error) {
	query := fmt.Sprintf("update %s set %s where `is_delete` = 0 and `uuid` = ?", m.table, objectUpdateFields)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	now := time.Now().Local()
	req, err :=stmt.ExecCtx(ctx, newData.ObjectName, newData.Identifier, newData.Key, newData.Sort, newData.Type, newData.SubType, newData.Extra, newData.Icon, newData.Status, newData.Puuid, now, newData.Uuid)
	if err !=nil{
		return 
	}
	
	rowsAffected, err := req.RowsAffected()
	if err !=nil{
		return
	}

	if rowsAffected == 0 {
		err = errors.New("no rows affected")
	}

	return err
}


// 用于首页展示统计数量
func (m *defaultObjectModel) CountByType(ctx context.Context) (resp *[]ObjectTypeCountItem, err error) {
	query := fmt.Sprintf("select `type`, count(*) as count from %s where `is_delete` = 0 and `status` =1 group by type", m.table)
	
	// 确定sql无参数，不需要Prepare
	resp = new([]ObjectTypeCountItem)
	err = m.conn.QueryRowsCtx(ctx, resp, query)

	return resp, err
}
