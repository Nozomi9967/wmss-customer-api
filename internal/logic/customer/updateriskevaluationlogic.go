// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRiskEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新风险测评
func NewUpdateRiskEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRiskEvaluationLogic {
	return &UpdateRiskEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRiskEvaluationLogic) UpdateRiskEvaluation(req *types.UpdateRiskEvaluationReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
