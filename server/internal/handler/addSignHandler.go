package handler

import (
	"net/http"

	"seal/server/internal/logic"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddSignHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddSignReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddSignLogic(r.Context(), svcCtx)
		resp, err := l.AddSign(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
