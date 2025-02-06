package svc

import (
	"demo/internal/config"
	"plugins/utils/verfy"
)

type ServiceContext struct {
	Config config.Config
	Verify *verfy.Verify
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Verify: verfy.NewVerify(),
	}
}
