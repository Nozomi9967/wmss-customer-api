// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
