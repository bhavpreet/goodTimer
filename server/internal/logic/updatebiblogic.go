package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBibLogic {
	return &UpdateBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBibLogic) UpdateBib(req *types.UpdateBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Bib)
	l.svcCtx.Get(req.ID, resp)

	if req.BibNo != "" {
		resp.No = req.BibNo
	}
	resp.StartTime = req.StartTime
	resp.EndTime = req.EndTime

	err = l.svcCtx.Update(req.ID, resp)
	if err != nil {
		return nil, err
	}

	resp = new(types.Bib)
	err = l.svcCtx.Get(req.ID, resp)
	return resp, err
}
