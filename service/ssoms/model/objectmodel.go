package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ObjectModel = (*customObjectModel)(nil)

type (
	// ObjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customObjectModel.
	ObjectModel interface {
		objectModel
	}

	customObjectModel struct {
		*defaultObjectModel
	}
)

// NewObjectModel returns a model for the database table.
func NewObjectModel(conn sqlx.SqlConn) ObjectModel {
	return &customObjectModel{
		defaultObjectModel: newObjectModel(conn),
	}
}
