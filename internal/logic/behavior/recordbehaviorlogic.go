// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordBehaviorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 记录客户行为
func NewRecordBehaviorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordBehaviorLogic {
	return &RecordBehaviorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordBehaviorLogic) RecordBehavior(req *types.RecordBehaviorReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
