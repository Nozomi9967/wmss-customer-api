// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package customer

import (
	"net/http"

	"github.com/Nozomi9967/wmss-customer-api/internal/logic/customer"
	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 申购风险测评
func RiskValidateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RiskValidityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewRiskValidateLogic(r.Context(), svcCtx)
		resp, err := l.RiskValidate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
