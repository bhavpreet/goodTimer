package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBibLogic {
	return &GetBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBibLogic) GetBib(req *types.GetBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line

	return
}
