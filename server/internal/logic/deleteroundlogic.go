package logic

import (
	"context"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoundLogic {
	return &DeleteRoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoundLogic) DeleteRound(req *types.DeleteRoundRequest) error {
	// todo: add your logic here and delete this line

	c, err := l.svcCtx.DB.GetCollection(l.ctx, req.Round)
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return err
	}

	err = c.DeleteCollection(l.ctx)
	if err != nil {
		logx.Errorf("Unable to delete, err: %v", err)
		return err
	}

	return err
}
