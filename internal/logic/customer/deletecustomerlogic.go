// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 逻辑删除客户信息
func NewDeleteCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerLogic {
	return &DeleteCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCustomerLogic) DeleteCustomer(req *types.DeleteCustomerReq) (resp *types.Response, err error) {
	err = l.svcCtx.CustomerInfoModel.DeleteLogical(l.ctx, req)
	if err != nil {
		l.Logger.Errorf("删除客户失败:%v", err)
		return &types.Response{
			Code: 400,
			Msg:  "删除客户失败",
		}, err
	}
	return &types.Response{
		Code: 200,
		Msg:  "删除客户成功",
	}, err
}
