// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/Nozomi9967/wmss-customer-api/utils"
	"github.com/Nozomi9967/wmss-user-api/common"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"

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
	userId, _ := l.ctx.Value("user_id").(string)
	requestUserInfo := &userRpc.GetUserRequest{
		UserId: userId,
	}
	user, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, requestUserInfo)
	if user == nil || err != nil {
		l.Logger.Errorf("用户[%s]不存在", userId)
		return &types.Response{
			Code: 404,
			Msg:  "用户不存在",
		}, nil
	}

	if user.RoleId != common.SUPER_ADMIN_ROLE_ID {
		l.Logger.Errorf("用户[%s]权限不足", userId)
		return &types.Response{
			Code: 403,
			Msg:  "权限不足，仅超级管理员可获取用户信息",
		}, nil
	}

	var customers []*model.CustomerInfo
	customers, err = l.svcCtx.CustomerInfoModel.FindBatches(l.ctx, req)
	if err != nil {
		l.Logger.Errorf("查询失败：%v", err)
		return &types.Response{
			Code: 400,
			Msg:  "查询失败",
		}, nil
	}

	//类型转换
	var customersInfo []*types.CustomerInfo
	for _, customer := range customers {
		var customerInfo *types.CustomerInfo
		customerInfo, err = utils.CustomerToCustomerInfo(customer)
		if err != nil {
			l.Logger.Errorf("查询失败：%v", err)
			return &types.Response{
				Code: 400,
				Msg:  "查询失败",
			}, nil
		}
		customersInfo = append(customersInfo, customerInfo)
	}
	return &types.Response{
		Code: 200,
		Msg:  "查询成功",
		Data: customersInfo,
	}, nil
}
