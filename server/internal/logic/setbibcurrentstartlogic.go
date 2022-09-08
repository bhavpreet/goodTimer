package logic

import (
	"context"
	"fmt"

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

	if bib.StartTime == "" {
		current, _ := GetCurrent(l.svcCtx.Store)
		current.CurrentRound = req.Round
		current.CurrentStartBib = req.Bib
		l.svcCtx.Store.Update("current", current)
	} else {
		return fmt.Errorf("Bib start time is set to %+v", bib)
	}

	return nil
}
