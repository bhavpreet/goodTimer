package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoundLogic {
	return &AddRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoundLogic) AddRound(req *types.AddRoundRequest) (resp *types.Round, err error) {
	// todo: add your logic here and delete this line

	// uid := uuid.NewString()

	r := &types.Round{
		// ID:     uid,
		Name:   req.Name,
		Status: "NONE",
	}

	_, err = l.svcCtx.DB.GetCollection(l.ctx, r.Name)
	if err != nil {
		logx.Errorf("Error while getting/creating colecltion (%s), err:%v", r.Name, err)
		return nil, err
	}

	// check if there is a current round, if not then set
	currentRound, err := getCurrentRound(l.ctx, l.svcCtx.DB)
	if err != nil || currentRound.CurrentRound == "" {
		scl := &SetRoundCurrentLogic{ctx: l.ctx, svcCtx: l.svcCtx}
		err = scl.SetRoundCurrent(&types.SetRoundCurrentRequest{Round: r.Name})
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
