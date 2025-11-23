package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerBehaviorModel = (*customCustomerBehaviorModel)(nil)

type (
	// CustomerBehaviorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerBehaviorModel.
	CustomerBehaviorModel interface {
		customerBehaviorModel
		withSession(session sqlx.Session) CustomerBehaviorModel
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
