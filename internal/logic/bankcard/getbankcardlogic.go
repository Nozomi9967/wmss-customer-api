// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/Nozomi9967/wmss-customer-api/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBankCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询银行卡详情
func NewGetBankCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBankCardLogic {
	return &GetBankCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBankCardLogic) GetBankCard(req *types.GetBankCardReq) (resp *types.Response, err error) {
	var bankCard *model.CustomerBankCard
	bankCard, err = l.svcCtx.CustomerBankCardModel.FindOne(l.ctx, req.CardId)
	if err != nil {
		l.Logger.Errorf("获取银行卡信息失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "获取银行卡信息失败",
		}, nil
	}

	var bankCardInfo *types.BankCardInfo
	bankCardInfo, err = utils.BankCardToBankCardInfo(bankCard)
	if err != nil {
		l.Logger.Errorf("获取银行卡信息失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "获取银行卡信息失败",
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "获取银行卡信息成功",
		Data: bankCardInfo,
	}, nil
}
