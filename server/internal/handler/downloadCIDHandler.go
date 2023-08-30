package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"seal/server/internal/logic"
	"seal/server/internal/svc"
	"seal/server/internal/types"
)

func DownloadCIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadCidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDownloadCIDLogic(r.Context(), svcCtx)
		resp, err := l.DownloadCID(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
