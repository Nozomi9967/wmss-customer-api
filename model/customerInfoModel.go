package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerInfoModel = (*customCustomerInfoModel)(nil)

type (
	// CustomerInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerInfoModel.
	CustomerInfoModel interface {
		customerInfoModel
		withSession(session sqlx.Session) CustomerInfoModel
		FindBatches(context context.Context, req *types.ListCustomerReq) ([]*CustomerInfo, error)
		DeleteLogical(context context.Context, req *types.DeleteCustomerReq) error
	}

	customCustomerInfoModel struct {
		*defaultCustomerInfoModel
	}
)

// NewCustomerInfoModel returns a model for the database table.
func NewCustomerInfoModel(conn sqlx.SqlConn) CustomerInfoModel {
	return &customCustomerInfoModel{
		defaultCustomerInfoModel: newCustomerInfoModel(conn),
	}
}

func (m *customCustomerInfoModel) withSession(session sqlx.Session) CustomerInfoModel {
	return NewCustomerInfoModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCustomerInfoModel) FindBatches(context context.Context, req *types.ListCustomerReq) ([]*CustomerInfo, error) {
	// 1. 构造 WHERE 子句和参数列表
	var whereBuilder strings.Builder
	var args []interface{}

	// 默认只查询未删除的记录
	whereBuilder.WriteString("WHERE deleted_at IS NULL")

	// 添加条件
	if req.CustomerName != "" {
		// 客户名称(模糊查询)
		whereBuilder.WriteString(" AND customer_name LIKE ?")
		args = append(args, "%"+req.CustomerName+"%")
	}

	if req.CustomerType != "" {
		// 客户类型(精确查询)
		whereBuilder.WriteString(" AND customer_type = ?")
		args = append(args, req.CustomerType)
	}

	if req.RiskLevel != "" {
		// 风险等级(精确查询)
		whereBuilder.WriteString(" AND risk_level = ?")
		args = append(args, req.RiskLevel)
	}

	if req.IdNumber != "" {
		// 证件号码(精确查询)
		whereBuilder.WriteString(" AND id_number = ?")
		args = append(args, req.IdNumber)
	}

	// 2. 构造 SQL 语句
	// 选择所有字段（注意：实际应用中推荐只选择需要的字段）
	query := fmt.Sprintf("SELECT %s FROM %s %s ORDER BY create_time DESC LIMIT ?, ?",
		m.columns(), // 假设有一个方法返回所有字段名
		m.table,     // 表名
		whereBuilder.String(),
	)

	// 3. 计算分页偏移量
	offset := (req.Page - 1) * req.PageSize

	// 添加分页参数
	args = append(args, offset)
	args = append(args, req.PageSize)

	// 4. 执行查询
	var customers []*CustomerInfo
	// 使用 sqlx 的 QueryRowsCtx 方法执行查询
	err := m.conn.QueryRowsCtx(context, &customers, query, args...)

	if err != nil && err != sqlx.ErrNotFound {
		return nil, err
	}

	// 如果没有找到记录，返回空列表而不是 nil，这通常更友好
	if customers == nil {
		return make([]*CustomerInfo, 0), nil
	}

	return customers, nil
}

// 辅助方法：返回表中的所有列名
// 在实际项目中，这个方法通常在 model 结构体初始化时定义好
func (m *customCustomerInfoModel) columns() string {
	// 这里只列出 CustomerInfo 中带 `db:"..."` tag 的字段
	return "customer_id, customer_name, customer_type, id_type, id_number, risk_level, risk_evaluation_time, risk_evaluation_expire_time, contact_phone, email, create_time, update_time, deleted_at"
}

func (m *customCustomerInfoModel) DeleteLogical(ctx context.Context, req *types.DeleteCustomerReq) error {
	// 1. 定义 SQL 语句：更新 deleted_at 字段
	// 假设您的表名为 'customer_info'，主键为 'customer_id'
	query := `UPDATE customer_info SET deleted_at = ? WHERE customer_id = ? AND deleted_at IS NULL`

	// 2. 获取当前时间
	// 逻辑删除通常需要精确的时间戳
	now := time.Now()

	// 3. 执行 SQL 语句
	_, err := m.conn.ExecCtx(ctx, query, now, req.CustomerId)

	if err != nil {
		// 记录错误日志
		logx.WithContext(ctx).Errorf("DeleteLogical error: customerId=%s, error=%v", req.CustomerId, err)
		return err
	}

	// 4. (可选) 清理缓存
	// 如果您的模型使用了缓存，您需要根据 Key 清理缓存
	// key := fmt.Sprintf("cache:customer_info:%s", req.CustomerId)
	// m.delCache(key)

	return nil
}
