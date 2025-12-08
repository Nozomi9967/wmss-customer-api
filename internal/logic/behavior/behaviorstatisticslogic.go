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
	// 时间格式
	layout := "2006-01-02 15:04:05"

	// 解析开始时间
	var startTime time.Time
	var sTime int64
	if req.StartTime != "" {
		startTime, err = time.Parse(layout, req.StartTime)
		if err != nil {
			l.Logger.Errorf("解析开始时间失败: %v", err)
			return &types.Response{
				Code: 400,
				Msg:  "开始时间格式错误，正确格式: 2006-01-02 15:04:05",
			}, nil
		}
		sTime = startTime.Unix()
	}

	// 解析结束时间
	var endTime time.Time
	var eTime int64
	if req.EndTime != "" {
		endTime, err = time.Parse(layout, req.EndTime)
		if err != nil {
			l.Logger.Errorf("解析结束时间失败: %v", err)
			return &types.Response{
				Code: 400,
				Msg:  "结束时间格式错误，正确格式: 2006-01-02 15:04:05",
			}, nil
		}
		eTime = endTime.Unix()
	}

	// 时间范围校验
	if req.StartTime != "" && req.EndTime != "" && sTime > eTime {
		return &types.Response{
			Code: 400,
			Msg:  "开始时间不能大于结束时间",
		}, nil
	}

	// 调用 RPC 查询行为列表（不分页，获取所有数据用于统计）
	reqRpc := customerRpc.ListBehaviorReq{
		CustomerId: req.CustomerId,
		StartTime:  sTime,
		EndTime:    eTime,
		Page:       1,
		PageSize:   10000, // 设置一个较大的值获取所有数据，或者多次调用分页获取
	}

	list, err := l.svcCtx.CustomerRpcClient.ListBehavior(l.ctx, &reqRpc)
	if err != nil {
		l.Logger.Errorf("查询行为列表失败: %v", err)
		return &types.Response{
			Code: 500,
			Msg:  "查询行为数据失败",
		}, nil
	}

	// 统计总数
	totalCount := list.Total

	// 按行为类型统计
	typeCountMap := make(map[string]int64)
	for _, item := range list.List {
		typeCountMap[item.BehaviorType]++
	}

	// 转换为响应格式
	typeStatistics := make([]types.BehaviorTypeCount, 0, len(typeCountMap))
	for behaviorType, count := range typeCountMap {
		typeStatistics = append(typeStatistics, types.BehaviorTypeCount{
			BehaviorType: behaviorType,
			Count:        count,
		})
	}

	// 构建响应
	respData := types.BehaviorStatisticsResp{
		TotalCount:     totalCount,
		TypeStatistics: typeStatistics,
	}

	return &types.Response{
		Code: 200,
		Msg:  "统计成功",
		Data: respData,
	}, nil
}
