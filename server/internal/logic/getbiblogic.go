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

func getBib(ctx context.Context,
	svcCtx *svc.ServiceContext,
	req *types.GetBibRequest) (resp *types.Bib, err error) {

	resp = new(types.Bib)
	err = svcCtx.Get(req.Bib, resp)
	if err != nil {
		return nil, err
	}

	r, err := NewGetRoundLogic(ctx, svcCtx).GetRound(&types.GetRoundRequest{
		ID: req.Round,
	})
	if err != nil {
		logx.Errorf("Did not find the Round, %s", req.Round)
		return nil, err
	}
	resp.Round = *r

	current, _ := GetCurrent(svcCtx.Store)
	if current.CurrentStartBib == resp.ID {
		resp.ISCurrentStart = true
	}
	if current.CurrentEndBib == resp.ID {
		resp.ISCurrentEnd = true
	}

	return resp, nil
}

func (l *GetBibLogic) GetBib(req *types.GetBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line
	resp, err = getBib(l.ctx, l.svcCtx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
