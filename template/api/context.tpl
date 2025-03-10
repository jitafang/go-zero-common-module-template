package svc

import (
	{{.configImport}}
	"plugins/utils/verfy"
)

type ServiceContext struct {
	Config {{.config}}
	Verify *verfy.Verify
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Verify: verfy.NewVerify(),
		{{.middlewareAssignment}}

	}
}
