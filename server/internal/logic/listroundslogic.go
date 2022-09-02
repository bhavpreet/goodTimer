package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoundsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoundsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoundsLogic {
	return &ListRoundsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoundsLogic) ListRounds() (resp *types.ListRoundsResp, err error) {
	// todo: add your logic here and delete this line
	c, err := l.svcCtx.DB.ListCollections(l.ctx)
	if err != nil {
		logx.Errorf("Error occured while listing collecitons, err :%v", err)
		return nil, err
	}

	currentRound, err := getCurrentRound(l.ctx, l.svcCtx.DB)
	if err != nil {
		logx.Errorf("Error occurred getCurrentRound")
		return nil, err
	}

	resp = new(types.ListRoundsResp)
	resp.Rounds = []types.Round{}

	for _, _c := range c {
		isCurrent := false

		if _c == "__current" {
			continue
		}
		if _c == currentRound.CurrentRound {
			isCurrent = true
		}
		resp.Rounds = append(resp.Rounds, types.Round{Name: _c, ISCurrent: isCurrent})
	}
	return
}
