package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentLogic {
	return &GetCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentLogic) GetCurrent() (resp *types.Current, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Current)
	l.svcCtx.Get("current", resp)
	return
}
