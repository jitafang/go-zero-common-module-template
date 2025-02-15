package {{.PkgName}}

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
    "plugins/entity/response"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		err := svcCtx.Verify.ValidateAndHandleError(req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerError(err.Error()))
			return
		}
		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		resp, err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerError(err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.HandlerSuccess(resp.Data))
		}
	}
}