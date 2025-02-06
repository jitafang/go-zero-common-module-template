package handler

import (
	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"plugins/entity/response"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		err := svcCtx.Verify.ValidateAndHandleError(req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerError(err.Error()))
			return
		}
		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerError(err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerSuccess(resp.Data))
		}
	}
}
