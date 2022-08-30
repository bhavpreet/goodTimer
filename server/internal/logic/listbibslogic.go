package logic

import (
	"context"
	"encoding/json"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBibsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBibsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBibsLogic {
	return &ListBibsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBibsLogic) ListBibs(req *types.ListBibsRequest) (resp []types.Bib, err error) {
	// todo: add your logic here and delete this line
	c, err := l.svcCtx.DB.GetCollection(l.ctx, req.Round)
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return nil, err
	}
	ret, err := c.List(l.ctx)
	if err != nil {
		logx.Errorf("Unable to list collection name %v, err = %v", req.Round, err)
		return nil, err
	}

	bibs := []types.Bib{}
	for _, r := range ret {
		bib := &types.Bib{}
		err = json.Unmarshal(r, bib)
		if err != nil {
			logx.Errorf("Unable to unmarshal bib")
			return nil, err
		}

		bibs = append(bibs, *bib)
	}
	return bibs, nil
}
