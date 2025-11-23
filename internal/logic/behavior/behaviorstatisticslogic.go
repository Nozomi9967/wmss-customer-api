// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BehaviorStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户行为统计
func NewBehaviorStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BehaviorStatisticsLogic {
	return &BehaviorStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BehaviorStatisticsLogic) BehaviorStatistics(req *types.BehaviorStatisticsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
