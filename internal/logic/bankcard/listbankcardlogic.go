// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"
	"fmt"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/Nozomi9967/wmss-customer-api/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBankCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户银行卡列表
func NewListBankCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBankCardLogic {
	return &ListBankCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBankCardLogic) ListBankCard(req *types.ListBankCardReq) (resp *types.Response, err error) {

	fmt.Print(req)

	var bankCards []*model.CustomerBankCard
	bankCards, err = l.svcCtx.CustomerBankCardModel.FindBatches(l.ctx, req)
	if err != nil {
		l.Logger.Errorf("获取银行卡信息失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "获取银行卡信息失败",
		}, nil
	}

	var bankCardsInfo []*types.BankCardInfo
	for _, bankCard := range bankCards {
		var bankCardInfo *types.BankCardInfo
		bankCardInfo, err = utils.BankCardToBankCardInfo(bankCard)
		if err != nil {
			l.Logger.Errorf("获取银行卡信息失败: %v", err)
			return &types.Response{
				Code: 400,
				Msg:  "获取银行卡信息失败",
			}, nil
		}
		bankCardsInfo = append(bankCardsInfo, bankCardInfo)
	}
	return &types.Response{
		Code: 200,
		Msg:  "获取银行卡信息成功",
		Data: bankCardsInfo,
	}, nil
}
