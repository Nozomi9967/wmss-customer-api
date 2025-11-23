// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"net/http"

	"WMSS/customer/api/internal/logic/behavior"
	"WMSS/customer/api/internal/svc"
	"WMSS/customer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 记录客户行为
func RecordBehaviorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordBehaviorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := behavior.NewRecordBehaviorLogic(r.Context(), svcCtx)
		resp, err := l.RecordBehavior(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
