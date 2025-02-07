package test

import (
	"demo/internal/logic/test"
	"demo/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"plugins/entity"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get()
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerError(err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerSuccess(resp.Data))
		}
	}
}
