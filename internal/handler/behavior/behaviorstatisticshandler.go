// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package behavior

import (
	"net/http"

	"github.com/Nozomi9967/wmss-customer-api/internal/logic/behavior"
	"github.com/Nozomi9967/wmss-customer-api/internal/svc"
	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 客户行为统计
func BehaviorStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BehaviorStatisticsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := behavior.NewBehaviorStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.BehaviorStatistics(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
