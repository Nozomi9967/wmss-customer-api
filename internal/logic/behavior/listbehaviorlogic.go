// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"context"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	customerRpc "github.com/Nozomi9967/wmss-customer-rpc/pb"
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
	layout := "2006-01-02 15:04:05"
	var startTime time.Time
	var sTime int64
	if req.StartTime != "" {
		startTime, err = time.Parse(layout, req.StartTime)
		sTime = startTime.Unix()

	}
	var endTime time.Time
	var eTime int64
	if req.EndTime != "" {
		endTime, err = time.Parse(layout, req.EndTime)
		eTime = endTime.Unix()
	}
	if err != nil {
		l.Logger.Errorf("解析失败:", err)
		return &types.Response{
			Code: 400,
			Msg:  "系统内部错误",
		}, nil
	}
	reqRpc := customerRpc.ListBehaviorReq{
		CustomerId:       req.CustomerId,
		BehaviorType:     req.BehaviorType,
		RelatedProductId: req.RelatedProductId,
		StartTime:        sTime,
		EndTime:          eTime,
		Page:             int32(req.Page),
		PageSize:         int32(req.PageSize),
	}
	var list *customerRpc.ListBehaviorResp
	list, err = l.svcCtx.CustomerRpcClient.ListBehavior(l.ctx, &reqRpc)
	if err != nil {
		l.Logger.Errorf("查询失败:", err)
		return &types.Response{
			Code: 400,
			Msg:  "系统内部错误",
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "查询成功",
		Data: list,
	}, nil
}
