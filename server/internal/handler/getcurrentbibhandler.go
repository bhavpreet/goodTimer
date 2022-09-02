package handler

import (
	"net/http"

	"github.com/bhavpreet/goodTimer/server/internal/logic"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCurrentBibHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCurrentBibReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetCurrentBibLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentBib(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
