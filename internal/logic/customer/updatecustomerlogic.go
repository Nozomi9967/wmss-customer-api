// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/Nozomi9967/wmss-user-api/common"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新客户信息
func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(req *types.UpdateCustomerReq) (resp *types.Response, err error) {
	// 权限校验
	userId, _ := l.ctx.Value("user_id").(string)
	requestUserInfo := &userRpc.GetUserRequest{
		UserId: userId,
	}
	user, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, requestUserInfo)
	if user == nil || err != nil {
		l.Logger.Errorf("用户[%s]不存在: %v", userId, err)
		return &types.Response{
			Code: 404,
			Msg:  "用户不存在",
		}, nil
	}

	if user.RoleId != common.SUPER_ADMIN_ROLE_ID {
		l.Logger.Errorf("用户[%s]权限不足", userId)
		return &types.Response{
			Code: 403,
			Msg:  "权限不足，仅超级管理员可更新客户信息",
		}, nil
	}

	// --- 核心更新逻辑开始 ---

	// 1. 根据 CustomerId 查询现有客户信息
	customerInfo, err := l.svcCtx.CustomerInfoModel.FindOne(l.ctx, req.CustomerId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("客户ID[%s]不存在", req.CustomerId)
			return &types.Response{
				Code: 404,
				Msg:  "客户信息不存在",
			}, nil
		}
		l.Logger.Errorf("查询客户信息失败, ID[%s]: %v", req.CustomerId, err)
		return nil, errors.Wrap(err, "query customer info failed") // 返回内部错误
	}

	// 2. 根据 req 中的字段进行选择性更新
	hasUpdated := false // 标记是否有字段被更新

	// CustomerName
	if req.CustomerName != "" {
		customerInfo.CustomerName = req.CustomerName
		hasUpdated = true
	}

	// RiskLevel
	if req.RiskLevel != "" {
		customerInfo.RiskLevel = req.RiskLevel
		hasUpdated = true
	}

	// ContactPhone (使用 sql.NullString)
	if req.ContactPhone != "" {
		customerInfo.ContactPhone = sql.NullString{String: req.ContactPhone, Valid: true}
		hasUpdated = true
	} else if req.ContactPhone == "" && customerInfo.ContactPhone.Valid {
		// 如果请求明确传空字符串，可以选择清空数据库中的值 (视业务需求而定)
		customerInfo.ContactPhone = sql.NullString{Valid: false}
		hasUpdated = true
	}

	// Email (使用 sql.NullString)
	if req.Email != "" {
		customerInfo.Email = sql.NullString{String: req.Email, Valid: true}
		hasUpdated = true
	} else if req.Email == "" && customerInfo.Email.Valid {
		// 如果请求明确传空字符串，可以选择清空数据库中的值 (视业务需求而定)
		customerInfo.Email = sql.NullString{Valid: false}
		hasUpdated = true
	}

	// RiskEvaluationTime (需要解析时间字符串)
	if req.RiskEvaluationTime != "" {
		t, timeErr := time.ParseInLocation("2006-01-02 15:04:05", req.RiskEvaluationTime, time.Local)
		if timeErr != nil {
			l.Logger.Errorf("风险测评时间格式错误: %s", req.RiskEvaluationTime)
			return &types.Response{
				Code: 400,
				Msg:  "风险测评时间格式错误，应为 YYYY-MM-DD HH:mm:ss",
			}, nil
		}
		customerInfo.RiskEvaluationTime = sql.NullTime{
			Time:  t,
			Valid: true,
		}
		hasUpdated = true
	}

	// RiskEvaluationExpireTime (需要解析时间字符串)
	if req.RiskEvaluationExpireTime != "" {
		t, timeErr := time.ParseInLocation("2006-01-02 15:04:05", req.RiskEvaluationExpireTime, time.Local)
		if timeErr != nil {
			l.Logger.Errorf("风险测评过期时间格式错误: %s", req.RiskEvaluationExpireTime)
			return &types.Response{
				Code: 400,
				Msg:  "风险测评过期时间格式错误，应为 YYYY-MM-DD HH:mm:ss",
			}, nil
		}
		customerInfo.RiskEvaluationExpireTime = sql.NullTime{
			Time:  t,
			Valid: true,
		}
		hasUpdated = true
	}

	// 3. 只有当有字段被更新时，才执行数据库操作
	if hasUpdated {
		// 设置 UpdateTime 为当前时间
		customerInfo.UpdateTime = time.Now()

		// 调用 model 的 Update 方法
		updateErr := l.svcCtx.CustomerInfoModel.Update(l.ctx, customerInfo)
		if updateErr != nil {
			l.Logger.Errorf("更新客户ID[%s]信息失败: %v", req.CustomerId, updateErr)
			return nil, errors.Wrap(updateErr, fmt.Sprintf("update customer ID[%s] failed", req.CustomerId))
		}
	} else {
		l.Logger.Infof("客户ID[%s]请求中没有可更新的字段", req.CustomerId)
	}

	// 4. 返回成功响应
	return &types.Response{
		Code: 200,
		Msg:  "客户信息更新成功",
	}, nil
}
