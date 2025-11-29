// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-user-api/common"
	userRpc "github.com/Nozomi9967/wmss-user-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type JwtAuthMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewJwtAuthMiddleware(svcCtx *svc.ServiceContext) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		svcCtx: svcCtx,
	}
}

func writeErrorResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// 使用 types.Response 结构体构造响应体
	respBody := types.Response{
		Code: uint32(code),
		Msg:  msg,
	}

	json.NewEncoder(w).Encode(respBody)
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 从请求的 Context 中获取 user_id。Go-Zero 已经自动存入。
		ctx := r.Context()

		// 使用 Context 和 ServiceContext
		l := logx.WithContext(ctx) // 获取日志记录器

		// 获取 userId
		userIdVal := ctx.Value("user_id")
		if userIdVal == nil {
			// 理论上如果配置了 jwt:Auth，这里不应该发生，除非 token 无效。
			writeErrorResponse(w, http.StatusUnauthorized, "未提供或无效的认证信息")
			return
		}
		userId := userIdVal.(string)

		// 权限校验
		requestUserInfo := &userRpc.GetUserRequest{
			UserId: userId,
		}

		// 注意：RPC 调用需要使用 context.Context
		user, err := m.svcCtx.UserRpcClient.GetUserInfo(ctx, requestUserInfo)

		if user == nil || err != nil {
			l.Errorf("RPC 调用或用户[%s]不存在: %v", userId, err)
			// 返回 401 或 404
			writeErrorResponse(w, http.StatusUnauthorized, "认证用户不存在或服务错误")
			return
		}

		if user.RoleId != common.SUPER_ADMIN_ROLE_ID {
			l.Errorf("用户[%s]权限不足, RoleID: %s", userId, user.RoleId)
			writeErrorResponse(w, http.StatusForbidden, "权限不足，仅超级管理员可操作此资源")
			return
		}

		// 权限校验通过，继续处理下一个 Handler
		next(w, r)
	}
}
