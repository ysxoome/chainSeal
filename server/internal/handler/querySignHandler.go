package handler

import (
	"net/http"

	"seal/server/internal/logic"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QuerySignHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySignReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewQuerySignLogic(r.Context(), svcCtx)
		resp, err := l.QuerySign(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
