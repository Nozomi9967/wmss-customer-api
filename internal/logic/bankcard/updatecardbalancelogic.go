// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"
	"strconv"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCardBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新银行卡余额
func NewUpdateCardBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCardBalanceLogic {
	return &UpdateCardBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCardBalanceLogic) UpdateCardBalance(req *types.UpdateCardBalanceReq) (resp *types.Response, err error) {
	var bankCard *model.CustomerBankCard
	bankCard, err = l.svcCtx.CustomerBankCardModel.FindOne(l.ctx, req.CardId)
	if err != nil {
		l.Logger.Errorf("更新银行卡余额失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "更新银行卡余额失败",
		}, nil
	}
	if req.Amount != "" {
		var amount float64
		amount, err = strconv.ParseFloat(req.Amount, 64)
		if err != nil {
			l.Logger.Errorf("更新银行卡余额失败: %v", err)
			return &types.Response{
				Code: 400,
				Msg:  "更新银行卡余额失败",
			}, nil
		}

		if req.OperateType == "consume" {
			bankCard.CardBalance -= amount
		} else {
			bankCard.CardBalance += amount
		}

	}

	err = l.svcCtx.CustomerBankCardModel.Update(l.ctx, bankCard)
	if err != nil {
		l.Logger.Errorf("更新银行卡余额失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "更新银行卡余额失败",
		}, nil
	}

	return &types.Response{
		Code: 200,
		Msg:  "更新银行卡余额成功",
	}, nil
}
