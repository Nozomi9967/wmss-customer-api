// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package bankcard

import (
	"net/http"

	"WMSS/customer/api/internal/logic/bankcard"
	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 解绑银行卡
func UnbindBankCardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnbindBankCardReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bankcard.NewUnbindBankCardLogic(r.Context(), svcCtx)
		resp, err := l.UnbindBankCard(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
