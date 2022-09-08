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

	var rounds []types.Round
	err  = l.svcCtx.Find(&rounds, nil)
	if err != nil {
		return nil, err
	}

	current, err := GetCurrent(l.svcCtx.Store)
	if err != nil {
		return nil, err
	}

	for i := range rounds {
		if rounds[i].ID == current.CurrentRound {
			rounds[i].ISCurrent = true
			break
		}
	}

	resp = new(types.ListRoundsResp)
	if rounds == nil {
		rounds = []types.Round{}
	}
	resp.Rounds = rounds
	return resp, nil
}
