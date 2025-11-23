// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户列表
func NewListCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCustomerLogic {
	return &ListCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCustomerLogic) ListCustomer(req *types.ListCustomerReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
