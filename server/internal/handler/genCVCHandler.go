package handler

import (
	"net/http"

	"seal/server/internal/logic"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenCVCHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CVCReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGenCVCLogic(r.Context(), svcCtx)
		resp, err := l.GenCVC(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
