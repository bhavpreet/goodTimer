package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/db"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentRoundLogic {
	return &GetCurrentRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentRoundLogic) GetCurrentRound() (resp string, err error) {
	// todo: add your logic here and delete this line
	return getCurrentRound(l.ctx, l.svcCtx.DB)
}

func getCurrentRound(ctx context.Context, db db.DB) (resp string, err error) {
	c, err := db.GetCollection(ctx, "__current")
	if err != nil {
		logx.Errorf("Unable to get collection name __current")
		return "", err
	}

	ret, err := c.Read(ctx, []byte("ROUND_CURRENT"))
	if err != nil {
		logx.Errorf("error reading current_round, err: %v", err)
		return "", err
	}

	logx.Infof("Current round is %s", string(ret))
	return string(ret), nil
}
