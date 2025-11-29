// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
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
	//userId, _ := l.ctx.Value("user_id").(string)
	//requestUserInfo := &userRpc.GetUserRequest{
	//	UserId: userId,
	//}
	//user, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, requestUserInfo)
	//if user == nil || err != nil {
	//	l.Logger.Errorf("用户[%s]不存在", userId)
	//	return &types.Response{
	//		Code: 404,
	//		Msg:  "用户不存在",
	//	}, nil
	//}
	//
	//if user.RoleId != common.SUPER_ADMIN_ROLE_ID {
	//	l.Logger.Errorf("用户[%s]权限不足", userId)
	//	return &types.Response{
	//		Code: 403,
	//		Msg:  "权限不足，仅超级管理员可获取用户信息",
	//	}, nil
	//}

	var customer *model.CustomerInfo
	customer, err = l.svcCtx.CustomerInfoModel.FindOne(l.ctx, req.CustomerId)
	if err != nil {
		l.Logger.Errorf("查询失败：%v", err)
		return &types.Response{
			Code: 400,
			Msg:  "查询失败",
		}, nil
	}

	// 类型转换
	var customerInfo *types.CustomerInfo
	customerInfo = &types.CustomerInfo{
		CustomerId:               customer.CustomerId,
		CustomerName:             customer.CustomerName,
		CustomerType:             customer.CustomerType,
		IdType:                   customer.IdType,
		IdNumber:                 customer.IdNumber,
		RiskLevel:                customer.RiskLevel,
		RiskEvaluationTime:       customer.RiskEvaluationTime.Format("2006-01-02 15:04:05"),
		RiskEvaluationExpireTime: customer.RiskEvaluationExpireTime.Format("2006-01-02 15:04:05"),
		ContactPhone:             customer.ContactPhone.String,
		Email:                    customer.Email.String,
		CreateTime:               customer.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:               customer.CreateTime.Format("2006-01-02 15:04:05"),
	}
	return &types.Response{
		Code: 200,
		Msg:  "查询成功",
		Data: customerInfo,
	}, nil
}
