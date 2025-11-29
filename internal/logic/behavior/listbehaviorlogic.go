// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBehaviorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询客户行为列表
func NewListBehaviorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBehaviorLogic {
	return &ListBehaviorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBehaviorLogic) ListBehavior(req *types.ListBehaviorReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
