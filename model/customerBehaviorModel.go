package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerBehaviorModel = (*customCustomerBehaviorModel)(nil)

type (
	// CustomerBehaviorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerBehaviorModel.
	CustomerBehaviorModel interface {
		customerBehaviorModel
		withSession(session sqlx.Session) CustomerBehaviorModel
		QueryRowNoCacheCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error
		QueryRowsNoCacheCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error
	}

	customCustomerBehaviorModel struct {
		*defaultCustomerBehaviorModel
	}
)

// NewCustomerBehaviorModel returns a model for the database table.
func NewCustomerBehaviorModel(conn sqlx.SqlConn) CustomerBehaviorModel {
	return &customCustomerBehaviorModel{
		defaultCustomerBehaviorModel: newCustomerBehaviorModel(conn),
	}
}

func (m *customCustomerBehaviorModel) withSession(session sqlx.Session) CustomerBehaviorModel {
	return NewCustomerBehaviorModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultCustomerBehaviorModel) QueryRowNoCacheCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return m.conn.QueryRowCtx(ctx, v, query, args...)
}

// QueryRowsNoCacheCtx 执行多行查询(不使用缓存)
func (m *defaultCustomerBehaviorModel) QueryRowsNoCacheCtx(ctx context.Context, v interface{}, query string, args ...interface{}) error {
	return m.conn.QueryRowsCtx(ctx, v, query, args...)
}
