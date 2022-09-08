package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoundLogic {
	return &GetRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoundLogic) GetRound(req *types.GetRoundRequest) (resp *types.Round, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Round)
	err = l.svcCtx.Get(req.ID, resp)
	if err != nil {
		return nil, err
	}
	current, err := GetCurrent(l.svcCtx.Store)
	if err != nil {
		return nil, err
	}

	if current.CurrentRound == resp.ID {
		resp.ISCurrent = true
	}
	return resp, err
}
