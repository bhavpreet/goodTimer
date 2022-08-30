package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBibCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetBibCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBibCurrentLogic {
	return &SetBibCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetBibCurrentLogic) SetBibCurrent(req *types.SetBibCurrentRequest) error {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Req = %+v", req)
	return nil
}
