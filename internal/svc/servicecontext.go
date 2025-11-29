package svc

import (
	"github.com/Nozomi9967/wmss-customer-api/internal/config"
	"github.com/Nozomi9967/wmss-customer-api/model"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UserRpcClient         userRpc.UserClient
	JwtAuthMiddleware     rest.Middleware
	CustomerBankCardModel model.CustomerBankCardModel `json:"customer_bank_card_model,omitempty"`
	CustomerBehaviorModel model.CustomerBehaviorModel `json:"customer_behavior_model,omitempty"`
	CustomerInfoModel     model.CustomerInfoModel     `json:"customer_info_model,omitempty"`
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		// 初始化 User RPC 客户端
		UserRpcClient: userRpc.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		// 注入 JWT 中间件
		//JwtAuthMiddleware:     middleware.JwtAuthMiddleware(c.Auth.AccessSecret),
		//JwtAuthMiddleware:     middleware.NewJwtAuthMiddleware().Handle,
		CustomerBankCardModel: model.NewCustomerBankCardModel(conn),
		CustomerBehaviorModel: model.NewCustomerBehaviorModel(conn),
		CustomerInfoModel:     model.NewCustomerInfoModel(conn),
	}
}
