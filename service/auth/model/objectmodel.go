package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ObjectModel = (*customObjectModel)(nil)

type (
	// ObjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customObjectModel.
	ObjectModel interface {
		objectModel
		RowBuilder() squirrel.SelectBuilder
		CustomFindOneByUuid(ctx context.Context, selectBuilder squirrel.SelectBuilder, uuid string) (resp *Object, err error)
		CustomFindOne(ctx context.Context, selectBuilder squirrel.SelectBuilder, args *ObjectFindOneArgs) (resp *Object, err error)
		FindMenusInRoleUUIDArray(ctx context.Context, topKey string, roleUUIDArray *[]string) (resp []*Menu, err error)
		FindObjectInRoleUUIDArray(ctx context.Context, topKey, key string, subType int64, roleUUIDArray *[]string) (exits bool, err error)
	}

	customObjectModel struct {
		*defaultObjectModel
	}
)

type Menu struct {
	Uuid       string `db:"uuid"`
	ObjectName string `db:"object_name"`
	Identifier string `db:"identifier"`
	Key        string `db:"key"`      // 操作对象的systemCode, 菜单的path, 操作的uri
	SubType    int64  `db:"sub_type"` // 子分类
	Icon       string `db:"icon"`     // 图标
	Puuid      string `db:"puuid"`    // 父级uuid
}

// NewObjectModel returns a model for the database table.
func NewObjectModel(conn sqlx.SqlConn, c cache.CacheConf) ObjectModel {
	return &customObjectModel{
		defaultObjectModel: newObjectModel(conn, c),
	}
}

// export logic
func (m *defaultObjectModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(objectRows).From(m.tableName())
}

func (m *defaultObjectModel) CustomFindOneByUuid(ctx context.Context, selectBuilder squirrel.SelectBuilder, uuid string) (resp *Object, err error) {
	query, values, err := selectBuilder.Where("uuid=?", uuid).ToSql()
	if err != nil {
		return
	}
	err = m.QueryRowNoCacheCtx(ctx, resp, query, values...)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}
	return
}

type ObjectFindOneArgs struct {
	Key     string
	TopKey  string
	Typ     int64
	SubType int64
}

func (m *defaultObjectModel) CustomFindOne(ctx context.Context, selectBuilder squirrel.SelectBuilder, args *ObjectFindOneArgs) (resp *Object, err error) {
	newSelectedBuilder := selectBuilder.Where("`status`=?", 1).Where("is_delete=?", 0)
	if args.TopKey == "" {
		err = errors.New("topKey is required")
		return
	}
	newSelectedBuilder = newSelectedBuilder.Where("top_key=?", args.TopKey)

	if args.Key != "" {
		newSelectedBuilder = newSelectedBuilder.Where("`key`=?", args.Key)
	}

	if args.Typ != 0 {
		newSelectedBuilder = newSelectedBuilder.Where("`type`=?", args.Typ)
	}

	if args.SubType != 0 {
		newSelectedBuilder = newSelectedBuilder.Where("sub_type=?", args.SubType)
	}

	query, values, err := newSelectedBuilder.ToSql()
	if err != nil {
		return
	}
	resp = new(Object)
	err = m.QueryRowNoCacheCtx(ctx, resp, query, values...)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}
	return
}

func (m *defaultObjectModel) FindMenusInRoleUUIDArray(ctx context.Context, topKey string, roleUUIDArray *[]string) (resp []*Menu, err error) {
	placeholder := []interface{}{topKey}
	fields := "o.`uuid`, o.`object_name`, o.`identifier`, o.`key`, o.`sub_type`, o.`icon`, o.`puuid`"
	where := "o.`is_delete` = 0 and o.`status` = 1 and o.`type` = 2 and p.`top_key` = ? and p.`is_delete` = 0"
	for i, roleUUID := range *roleUUIDArray {
		if i == 0 {
			where += " and p.`role_uuid` in("
		} else {
			where += ", "
		}
		where += "?"
		if i == len(*roleUUIDArray)-1 {
			where += ")"
		}
		placeholder = append(placeholder, roleUUID)
	}
	// 表名是否要用*defaultPermissionModel.tableName()来获取？
	query := fmt.Sprintf("select %s from %s as o left join permission as p on o.uuid = p.`object_uuid` where %s order by o.`sort`, o.`create_time`", fields, m.table, where)
	resp = make([]*Menu, 0, 1)
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, placeholder...)

	return resp, err
}

func (m *defaultObjectModel) FindObjectInRoleUUIDArray(ctx context.Context, topKey, key string, subType int64, roleUUIDArray *[]string) (exits bool, err error) {
	placeholder := []interface{}{topKey, key}
	where := "o.`is_delete` = 0 and o.`status` = 1 and o.`type` = 3 and p.`top_key` = ? and o.`key`=? and p.`is_delete` = 0"
	for i, roleUUID := range *roleUUIDArray {
		if i == 0 {
			where += " and p.`role_uuid` in("
		} else {
			where += ", "
		}
		where += "?"
		if i == len(*roleUUIDArray)-1 {
			where += ")"
		}
		placeholder = append(placeholder, roleUUID)
	}
	if subType != 0 {
		where += " and o.`sub_type` = ?"
		placeholder = append(placeholder, subType)
	}
	query := fmt.Sprintf("select count(*) from %s as o left join permission as p on o.uuid = p.`object_uuid` where %s limit 1", m.table, where)
	var had int
	err = m.QueryRowNoCacheCtx(ctx, &had, query, placeholder...)

	if err != nil {
		return false, err
	}

	if had == 0 {
		return false, ErrNotFound
	}

	return true, nil
}
