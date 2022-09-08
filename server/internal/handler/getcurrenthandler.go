package handler

import (
	"net/http"

	"github.com/bhavpreet/goodTimer/server/internal/logic"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCurrentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCurrentLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrent()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
