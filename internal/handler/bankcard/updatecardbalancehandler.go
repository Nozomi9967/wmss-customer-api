// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"net/http"

	"github.com/Nozomi9967/wmss-customer-api/internal/logic/bankcard"
	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新银行卡余额
func UpdateCardBalanceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCardBalanceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bankcard.NewUpdateCardBalanceLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCardBalance(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
