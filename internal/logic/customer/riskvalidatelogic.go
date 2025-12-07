// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RiskValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申购风险测评
func NewRiskValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RiskValidateLogic {
	return &RiskValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RiskValidateLogic) RiskValidate(req *types.RiskValidityReq) (resp *types.Response, err error) {
	// 1.
	var customer *model.CustomerInfo
	customer, err = l.svcCtx.CustomerInfoModel.FindOne(l.ctx, req.CustomerId)
	if err != nil {
		return &types.Response{
			Code: 400,
			Msg:  "申购风险校验失败",
			Data: &types.RiskValidityResp{
				IsValid: "false",
				Reason:  "客户不存在",
			},
		}, nil
	}

	// 2.
	riskList := []string{"R1", "R2", "R3", "R4", "R5"}
	flag := false
	for _, risk := range riskList {
		if customer.RiskLevel == risk {
			flag = true
			break
		}
	}
	if !flag {
		return &types.Response{
			Code: 400,
			Msg:  "申购风险校验失败",
			Data: &types.RiskValidityResp{
				IsValid: "false",
				Reason:  "客户风险等级不存在",
			},
		}, nil
	}

	// 3.
	if !customer.RiskEvaluationTime.Valid {
		return &types.Response{
			Code: 400,
			Msg:  "申购风险校验失败",
			Data: &types.RiskValidityResp{
				IsValid:   "false",
				Reason:    "用户未进行风险评估",
				RiskLevel: customer.RiskLevel,
			},
		}, nil
	}

	// 4.
	now := time.Now()
	if now.After(customer.RiskEvaluationExpireTime.Time) {
		return &types.Response{
			Code: 400,
			Msg:  "申购风险校验失败",
			Data: &types.RiskValidityResp{
				IsValid:    "false",
				Reason:     "用户风险评估过期",
				RiskLevel:  customer.RiskLevel,
				ExpireTime: customer.RiskEvaluationTime.Time.Format("2006-01-02 15:04:05"),
			},
		}, nil
	}

	return &types.Response{
		Code: 200,
		Msg:  "申购风险校验成功",
		Data: &types.RiskValidityResp{
			IsValid:    "true",
			RiskLevel:  customer.RiskLevel,
			ExpireTime: customer.RiskEvaluationTime.Time.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
