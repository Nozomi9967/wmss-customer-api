// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户详情
func NewGetCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerLogic {
	return &GetCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerLogic) GetCustomer(req *types.GetCustomerReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
