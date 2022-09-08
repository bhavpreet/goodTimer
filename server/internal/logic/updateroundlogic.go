package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoundLogic {
	return &UpdateRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoundLogic) UpdateRound(req *types.UpdateRoundRequest) (resp *types.Round, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Round)
	l.svcCtx.Get(req.ID, resp)

	if req.Name != "" {
		resp.Name = req.Name
	}

	err = l.svcCtx.Update(req.ID, resp)
	if err != nil {
		return nil, err
	}

	resp = new(types.Round)
	err = l.svcCtx.Get(req.ID, resp)
	return resp, err
}
