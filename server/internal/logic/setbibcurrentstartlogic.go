package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetBibCurrentStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetBibCurrentStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBibCurrentStartLogic {
	return &SetBibCurrentStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetBibCurrentStartLogic) SetBibCurrentStart(req *types.SetBibCurrentReq) error {
	// todo: add your logic here and delete this line

	bib, err := NewGetBibLogic(l.ctx, l.svcCtx).GetBib(&types.GetBibRequest{
		Round: req.Round,
		Bib:   req.Bib,
	})
	if err != nil {
		return err
	}


	current, _ := GetCurrent(l.svcCtx.Store)
	// toggle
	if bib.ID == current.CurrentStartBib { 
		current.CurrentStartBib = ""
	} else {
		current.CurrentRound = req.Round
		current.CurrentStartBib = bib.ID
	}

	l.svcCtx.Store.Update("current", current)

	return nil
}
