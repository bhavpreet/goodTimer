package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoundCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoundCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoundCurrentLogic {
	return &SetRoundCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoundCurrentLogic) SetRoundCurrent(req *types.SetRoundCurrentRequest) error {
	// todo: add your logic here and delete this line
	_, err := NewGetRoundLogic(l.ctx, l.svcCtx).GetRound(&types.GetRoundRequest{
		ID: req.Round,
	})
	if err != nil {
		return err
	}

	current, _ := GetCurrent(l.svcCtx.Store)
	current.CurrentRound = req.Round
	l.svcCtx.Store.Update("current", current)

	return nil
}
