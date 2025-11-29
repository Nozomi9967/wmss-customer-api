// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
	"github.com/Nozomi9967/wmss-user-api/common"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建客户
func NewCreateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCustomerLogic {
	return &CreateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomerLogic) CreateCustomer(req *types.CreateCustomerReq) (resp *types.Response, err error) {
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
			Msg:  "权限不足，仅超级管理员可新增权限",
		}, nil
	}

	var customer *model.CustomerInfo
	rawId := uuid.New().String()
	customerId := strings.ReplaceAll(rawId, "-", "")
	evalTime, err := time.Parse("2006-01-02 15:04:05", req.RiskEvaluationTime)
	if err != nil {
		// 解析失败：返回前端“时间格式错误”
		return &types.Response{
			Code: 400,
			Msg:  "风险测评时间格式错误，请使用 YYYY-MM-DD HH:MM:SS 格式（如 2025-11-28 14:30:00）",
		}, nil
	}
	expireTime, err := time.Parse("2006-01-02 15:04:05", req.RiskEvaluationExpireTime)
	if err != nil {
		return &types.Response{
			Code: 400,
			Msg:  "风险测评过期时间格式错误，请使用 YYYY-MM-DD HH:MM:SS 格式（如 2025-11-28 14:30:00）",
		}, nil
	}
	customer = &model.CustomerInfo{
		CustomerId:               customerId,
		CustomerName:             req.CustomerName,
		CustomerType:             req.CustomerType,
		IdType:                   req.IdType,
		IdNumber:                 req.IdNumber,
		RiskLevel:                req.RiskLevel,
		RiskEvaluationTime:       evalTime,
		RiskEvaluationExpireTime: expireTime,
		ContactPhone:             sql.NullString{String: req.ContactPhone, Valid: true},
		Email:                    sql.NullString{String: req.Email, Valid: true},
		CreateTime:               time.Now(),
		UpdateTime:               time.Now(),
		DeletedAt:                sql.NullTime{},
	}
	_, err = l.svcCtx.CustomerInfoModel.Insert(l.ctx, customer)
	if err != nil {
		l.Logger.Errorf("新增客户失败：%v", err)
		return &types.Response{
			Code: 400,
			Msg:  "新增失败",
		}, nil
	}

	return &types.Response{
		Code: 200,
		Msg:  "新建成功",
	}, nil
}
