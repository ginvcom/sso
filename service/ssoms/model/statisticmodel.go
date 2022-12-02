package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StatisticModel = (*customStatisticModel)(nil)

type (
	// StatisticModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatisticModel.
	StatisticModel interface {
		statisticModel
		FindYearsData(ctx context.Context) ([]*Statistic, error)
	}

	customStatisticModel struct {
		*defaultStatisticModel
	}
)

// NewStatisticModel returns a model for the database table.
func NewStatisticModel(conn sqlx.SqlConn) StatisticModel {
	return &customStatisticModel{
		defaultStatisticModel: newStatisticModel(conn),
	}
}

func (m *defaultStatisticModel) FindYearsData(ctx context.Context) ([]*Statistic, error) {
	query := fmt.Sprintf("select %s from %s order by id desc limit 12", statisticRows, m.table)

	var resp []*Statistic
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
