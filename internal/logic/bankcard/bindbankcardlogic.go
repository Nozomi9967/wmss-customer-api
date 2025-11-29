// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"
	"database/sql"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindBankCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定银行卡
func NewBindBankCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindBankCardLogic {
	return &BindBankCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindBankCardLogic) BindBankCard(req *types.BindBankCardReq) (resp *types.Response, err error) {

	var bankCard *model.CustomerBankCard
	now := time.Now()
	bankCard = &model.CustomerBankCard{
		CustomerId:     req.CustomerId,
		BankCardNumber: req.BankCardNumber,
		BankName:       req.BankName,
		CardBalance:    0,
		IsVirtual:      int64(req.IsVirtual),
		BindStatus:     "正常",
		BindTime:       now,
		UnbindTime:     sql.NullTime{},
		CreateTime:     now,
		UpdateTime:     now,
		DeletedAt:      sql.NullTime{},
	}
	_, err = l.svcCtx.CustomerBankCardModel.Insert(l.ctx, bankCard)
	if err != nil {
		l.Logger.Errorf("绑定银行卡失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "绑定银行卡失败",
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "绑定银行卡成功",
	}, nil
}
