package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBibLogic {
	return &AddBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBibLogic) AddBib(req *types.AddBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line

	round := new(types.Round)
	err = l.svcCtx.Get(req.Round, round)
	if err != nil {
		logx.Errorf("Unable to get round id %v", req.Round)
		return nil, err
	}

	// current, err := GetCurrent(l.svcCtx.Store)
	// if err != nil {
	// 	logx.Errorf("Unable to get current %v", err)
	// 	return nil, err
	// }

	logx.Infof("Round = %+v", round)
	bib := &types.Bib{
		ID:     uuid.NewString(),
		No:     req.Bib,
		Status: "NONE",
		ParentRoundID: round.ID,
	}

	logx.Infof("saving bib = %+v", bib)
	err = l.svcCtx.Insert(bib.ID, bib)
	if err != nil {
		logx.Errorf("Unable to save bib %+v, err: %v", bib, err)
		return nil, err
	}

	// if current.CurrentStartBib == "" {
	// 	SetCurrentStartBib(l.svcCtx.Store, bib.ID)
	// 	bib.ISCurrentStart = true
	// }

	// if current.CurrentEndBib == "" {
	// 	SetCurrentEndBib(l.svcCtx.Store, bib.ID)
	// 	bib.ISCurrentEnd = true
	// }

	l.svcCtx.Get(bib.ID, &bib)
	logx.Infof("Retrieved bib = %+v", bib)
	return bib, err
}
