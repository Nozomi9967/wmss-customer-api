// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
