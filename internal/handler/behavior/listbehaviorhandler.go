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

// 查询客户行为列表
func ListBehaviorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListBehaviorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := behavior.NewListBehaviorLogic(r.Context(), svcCtx)
		resp, err := l.ListBehavior(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
