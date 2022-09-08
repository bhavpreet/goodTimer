package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentRoundLogic {
	return &GetCurrentRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentRoundLogic) GetCurrentRound() (resp *types.GetCurrentRoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
