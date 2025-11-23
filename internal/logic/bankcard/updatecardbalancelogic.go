// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
