package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ObjectLogModel = (*customObjectLogModel)(nil)

type (
	// ObjectLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customObjectLogModel.
	ObjectLogModel interface {
		objectLogModel
	}

	customObjectLogModel struct {
		*defaultObjectLogModel
	}
)

// NewObjectLogModel returns a model for the database table.
func NewObjectLogModel(conn sqlx.SqlConn) ObjectLogModel {
	return &customObjectLogModel{
		defaultObjectLogModel: newObjectLogModel(conn),
	}
}
