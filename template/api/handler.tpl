package {{.PkgName}}

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		err := svcCtx.Verify.ValidateAndHandleError(req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: -1, Message: err.Error(), Data: nil})
			return
		}
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: -1,  Message: err.Error(), Data: nil})
		} else {
			httpx.OkJsonCtx(r.Context(), w, &response{Code: 0, Message: "success", {{if .HasResp}}Data: resp.Data{{end}}})
		}
	}
}