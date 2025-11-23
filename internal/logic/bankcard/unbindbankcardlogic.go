// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
