package logic

import (
	"context"
	"encoding/json"

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
	c, err := svcCtx.DB.GetCollection(ctx, req.Round)
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return nil, err
	}

	bib, err := c.Read(ctx, []byte(req.Bib))
	if err != nil {
		logx.Errorf(
			"Unable to read bib %s, from collection name %s, err: %v", req.Bib, req.Round, err)
		return nil, err
	}
	resp = &types.Bib{}
	err = json.Unmarshal(bib, resp)
	if err != nil {
		logx.Errorf("Unable to unmarshal bib")
		return nil, err
	}
	return resp, nil
}

func (l *GetBibLogic) GetBib(req *types.GetBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line
	resp, err = getBib(l.ctx, l.svcCtx, req)
	if err != nil {
		return nil, err
	}

	currentBib, err :=
		getCurrentBib(l.ctx, l.svcCtx, &types.GetCurrentBibReq{Round: req.Round})
	if err != nil {
		logx.Errorf("Unable to get current bib")
		return nil, err
	}

	if resp.ID == currentBib.CurrentBib {
		resp.ISCurrent = true
	}

	return resp, nil
}
