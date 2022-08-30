package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBibStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetBibStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBibStatusLogic {
	return &SetBibStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetBibStatusLogic) SetBibStatus(req *types.SetBibStatusRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
