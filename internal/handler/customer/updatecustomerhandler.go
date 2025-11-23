// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"net/http"

	"WMSS/customer/api/internal/logic/customer"
	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新客户信息
func UpdateCustomerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCustomerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewUpdateCustomerLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCustomer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
