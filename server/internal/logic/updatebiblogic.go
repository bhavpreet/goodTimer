package logic

import (
	"context"
	"encoding/json"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBibLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBibLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBibLogic {
	return &UpdateBibLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBibLogic) UpdateBib(req *types.UpdateBibRequest) (resp *types.Bib, err error) {
	// todo: add your logic here and delete this line

	getBibLogic := NewGetBibLogic(l.ctx, l.svcCtx)
	bib, err := getBibLogic.GetBib(&types.GetBibRequest{
		Round: req.Round,
		Bib:   req.Bib,
	})

	bib.ID = req.BibDoc.ID
	bib.StartTime = req.BibDoc.StartTime
	bib.EndTime = req.BibDoc.EndTime

	j, err := json.Marshal(bib)
	if err != nil {
		logx.Errorf("Unable to get json marshal %+v", bib)
		return nil, err
	}

	c, err := l.svcCtx.DB.GetCollection(l.ctx, req.Round)
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return nil, err
	}

	err = c.Write(l.ctx, []byte(req.Bib), j)
	if err != nil {
		logx.Errorf("Unable to Write %+v, err: %v", bib, err)
		return nil, err
	}

	return bib, ni
	l
}
