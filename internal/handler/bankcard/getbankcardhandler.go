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

// 查询银行卡详情
func GetBankCardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBankCardReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bankcard.NewGetBankCardLogic(r.Context(), svcCtx)
		resp, err := l.GetBankCard(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
