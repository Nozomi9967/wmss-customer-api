// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"
	"errors"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-user-api/common"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRiskEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新风险测评
func NewUpdateRiskEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRiskEvaluationLogic {
	return &UpdateRiskEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRiskEvaluationLogic) UpdateRiskEvaluation(req *types.UpdateRiskEvaluationReq) (resp *types.Response, err error) {

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

	// 1. 根据 CustomerId 查询现有客户信息
	customerInfo, dbErr := l.svcCtx.CustomerInfoModel.FindOne(l.ctx, req.CustomerId)
	if dbErr != nil {
		if errors.Is(dbErr, sqlx.ErrNotFound) { // 假设 model.ErrNotFound 就是 sqlx.ErrNotFound
			l.Logger.Errorf("客户ID[%s]不存在", req.CustomerId)
			return &types.Response{
				Code: 404,
				Msg:  "客户信息不存在",
			}, nil
		}
		l.Logger.Errorf("查询客户信息失败, ID[%s]: %v", req.CustomerId, dbErr)
		// 数据库查询错误，返回通用错误信息
		return &types.Response{
			Code: 500,
			Msg:  "系统繁忙，请稍后重试",
		}, nil
	}

	// 2. 解析并校验时间格式
	const timeLayout = "2006-01-02 15:04:05"

	// 解析 RiskEvaluationTime
	evalTime, timeErr := time.ParseInLocation(timeLayout, req.RiskEvaluationTime, time.Local)
	if timeErr != nil {
		l.Logger.Errorf("客户ID[%s]风险测评时间格式错误: %s", req.CustomerId, req.RiskEvaluationTime)
		return &types.Response{
			Code: 400,
			Msg:  "风险测评时间格式错误，应为 YYYY-MM-DD HH:mm:ss",
		}, nil
	}

	// 解析 RiskEvaluationExpireTime
	expireTime, timeErr := time.ParseInLocation(timeLayout, req.RiskEvaluationExpireTime, time.Local)
	if timeErr != nil {
		l.Logger.Errorf("客户ID[%s]风险测评过期时间格式错误: %s", req.CustomerId, req.RiskEvaluationExpireTime)
		return &types.Response{
			Code: 400,
			Msg:  "风险测评过期时间格式错误，应为 YYYY-MM-DD HH:mm:ss",
		}, nil
	}

	// 3. 更新客户信息字段

	// 检查是否有实际更新
	isUpdated := false

	// 仅更新必填的风险相关字段
	if customerInfo.RiskLevel != req.RiskLevel {
		customerInfo.RiskLevel = req.RiskLevel
		isUpdated = true
	}

	if !customerInfo.RiskEvaluationTime.Equal(evalTime) {
		customerInfo.RiskEvaluationTime = evalTime
		isUpdated = true
	}

	if !customerInfo.RiskEvaluationExpireTime.Equal(expireTime) {
		customerInfo.RiskEvaluationExpireTime = expireTime
		isUpdated = true
	}

	// 4. 执行数据库更新
	if isUpdated {
		// 更新 UpdateTime 字段
		customerInfo.UpdateTime = time.Now()

		updateErr := l.svcCtx.CustomerInfoModel.Update(l.ctx, customerInfo)
		if updateErr != nil {
			l.Logger.Errorf("更新客户ID[%s]的风险测评信息失败: %v", req.CustomerId, updateErr)
			// 数据库更新错误，返回通用错误信息
			return &types.Response{
				Code: 500,
				Msg:  "系统繁忙，风险测评更新失败",
			}, nil
		}
	} else {
		l.Logger.Infof("客户ID[%s]的风险测评信息没有变化，跳过数据库更新", req.CustomerId)
	}

	// 5. 返回成功响应
	return &types.Response{
		Code: 200,
		Msg:  "风险测评信息更新成功",
	}, nil
}
