package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBibLogic {
	return &DeleteBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBibLogic) DeleteBib() error {
	// todo: add your logic here and delete this line

	return nil
}
