package svc

import (
	"demo/internal/config"
	"demo/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"plugins/utils/verfy"
)

type ServiceContext struct {
	Config         config.Config
	Verify         *verfy.Verify
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Verify:         verfy.NewVerify(),
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}
