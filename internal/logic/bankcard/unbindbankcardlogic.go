// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindBankCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑银行卡
func NewUnbindBankCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindBankCardLogic {
	return &UnbindBankCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindBankCardLogic) UnbindBankCard(req *types.UnbindBankCardReq) (resp *types.Response, err error) {
	err = l.svcCtx.CustomerBankCardModel.DeleteLogical(l.ctx, req)
	if err != nil {
		l.Logger.Errorf("解绑银行卡失败: %v", err)
		return &types.Response{
			Code: 400,
			Msg:  "解绑银行卡失败",
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "解绑银行卡成功",
	}, nil
}
