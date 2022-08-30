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

	c, err := l.svcCtx.GetCollection(l.ctx, "__current")
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return err
	}

	logx.Infof("Writing current as %+v", req)
	err = c.Write(l.ctx, []byte("ROUND_CURRENT"), []byte(req.Round))
	if err != nil {
		logx.Errorf("Unable to write current round to %s", req.Round)
		return err
	}

	r, _ := c.Read(l.ctx, []byte("ROUND_CURRENT"))
	logx.Infof("Read %s", string(r))
	return nil
}
