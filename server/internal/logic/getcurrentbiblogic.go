package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentBibLogic {
	return &GetCurrentBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentBibLogic) GetCurrentBib(
	req *types.GetCurrentBibReq) (resp *types.GetCurrentBibResp, err error) {
	// todo: add your logic here and delete this line
	return getCurrentBib(l.ctx, l.svcCtx, req)
}

func getCurrentBib(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
	req *types.GetCurrentBibReq) (resp *types.GetCurrentBibResp, err error) {

	c, err := svcCtx.GetCollection(ctx, "__current")
	if err != nil {
		logx.Errorf("Unable to get collection name __current")
		return nil, err
	}

	currentBib, err := c.Read(ctx, []byte("CURRENT_BIB_" + req.Round))
	if err != nil {
		logx.Errorf("Unable to read from collection name %s", "CURRENT_BIB")
		return nil, err
	}

	resp = &types.GetCurrentBibResp{CurrentBib: string(currentBib)}
	return resp, nil
}
