package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerBankCardModel = (*customCustomerBankCardModel)(nil)

type (
	// CustomerBankCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerBankCardModel.
	CustomerBankCardModel interface {
		customerBankCardModel
		withSession(session sqlx.Session) CustomerBankCardModel
		FindBatches(context context.Context, req *types.ListBankCardReq) (resp []*CustomerBankCard, err error)
		DeleteLogical(context context.Context, req *types.UnbindBankCardReq) (err error)
	}

	customCustomerBankCardModel struct {
		*defaultCustomerBankCardModel
	}
)

// NewCustomerBankCardModel returns a model for the database table.
func NewCustomerBankCardModel(conn sqlx.SqlConn) CustomerBankCardModel {
	return &customCustomerBankCardModel{
		defaultCustomerBankCardModel: newCustomerBankCardModel(conn),
	}
}

func (m *customCustomerBankCardModel) DeleteLogical(ctx context.Context, req *types.UnbindBankCardReq) error {
	// 1. 获取当前时间
	now := time.Now()

	// 2. 定义更新 SQL
	// 我们同时更新以下字段：
	// - deleted_at: 标记为逻辑删除
	// - bind_status: 明确标记为已解绑
	// - unbind_time: 记录解绑的具体时间
	// - update_time: 记录最后修改时间
	// 假设表名为 `customer_bank_card` (这通常在 model 的 m.table 字段中)
	query := fmt.Sprintf("UPDATE %s SET `deleted_at` = ?, `bind_status` = ?, `unbind_time` = ?, `update_time` = ? WHERE `card_id` = ?", m.table)

	// 3. 定义状态常量
	const StatusUnbound = "已解绑"

	// 4. 执行更新
	// 注意：虽然 struct 中 DeletedAt 是 sql.NullTime，但 update 时直接传入 time.Time 即可
	result, err := m.conn.ExecCtx(ctx, query, now, StatusUnbound, now, now, req.CardId)
	if err != nil {
		logx.WithContext(ctx).Errorf("DeleteLogical exec error: %v, cardId: %d", err, req.CardId)
		return err
	}

	// 5. (可选) 检查是否真的更新了数据
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// 如果没有行受到影响，可能是 ID 不存在或者已经被删除了
		// 根据业务需求，这里可以返回 nil (视为幂等成功) 或者返回 ErrNotFound
		logx.WithContext(ctx).Infof("DeleteLogical: no rows affected, cardId: %d may not exist or already deleted", req.CardId)
		// return sqlx.ErrNotFound // 如果需要报错则取消注释
	}

	return nil
}

func (m *customCustomerBankCardModel) withSession(session sqlx.Session) CustomerBankCardModel {
	return NewCustomerBankCardModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCustomerBankCardModel) FindBatches(ctx context.Context, req *types.ListBankCardReq) (resp []*CustomerBankCard, err error) {

	// 1. 构建基础 SQL 查询
	var sb strings.Builder
	sb.WriteString("SELECT card_id, customer_id, bank_card_number, bank_name, card_balance, is_virtual, bind_status, bind_time, unbind_time, create_time, update_time, deleted_at FROM customer_bank_card WHERE ")

	// 2. 构建 WHERE 条件和参数列表
	var conditions []string
	var args []interface{}

	// a. 客户ID (必填条件)
	conditions = append(conditions, "customer_id = ?")
	args = append(args, req.CustomerId)

	// b. 绑定状态 (可选条件)
	if req.BindStatus != "" {
		conditions = append(conditions, "bind_status = ?")
		args = append(args, req.BindStatus)
	}

	// c. 是否虚拟卡 (可选条件)
	// 假设 IsVirtual 为 -1 时表示查询所有，0/1 时为精确查询
	if req.IsVirtual != -1 {
		conditions = append(conditions, "is_virtual = ?")
		args = append(args, req.IsVirtual)
	}

	// d. 排除逻辑删除的记录
	conditions = append(conditions, "deleted_at IS NULL")

	// 3. 拼接 WHERE 子句
	sb.WriteString(strings.Join(conditions, " AND "))

	// 4. (可选) 添加排序，通常按创建时间降序
	sb.WriteString(" ORDER BY create_time DESC")

	query := sb.String()

	// 5. 执行查询
	resp = make([]*CustomerBankCard, 0)
	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)

	if err != nil {
		if err == sqlx.ErrNotFound || err == sql.ErrNoRows {
			// 没有找到记录，返回空列表而不是错误
			return resp, nil
		}
		logx.WithContext(ctx).Errorf("FindBatches error: query=%s, args=%v, error=%v", query, args, err)
		return nil, err
	}

	return resp, nil
}
