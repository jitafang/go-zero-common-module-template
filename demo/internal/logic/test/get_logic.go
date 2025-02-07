package test

import (
	"context"

	"demo/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"plugins/entity/response"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp *response.LogicResponse, err error) {
	// todo: add your logic here and delete this line

	return response.LogicResult(nil), nil
}
