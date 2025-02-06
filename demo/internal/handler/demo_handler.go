package handler

import (
	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	type response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		err := svcCtx.Verify.ValidateAndHandleError(req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: -1, Message: err.Error(), Data: nil})
			return
		}
		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: -1, Message: err.Error(), Data: nil})
		} else {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: 0, Message: "success", Data: resp.Data})
		}
	}
}
