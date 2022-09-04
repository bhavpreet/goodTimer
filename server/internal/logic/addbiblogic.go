package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"

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

	c, err := l.svcCtx.DB.GetCollection(l.ctx, req.Round)
	if err != nil {
		logx.Errorf("Unable to get collection name %v", req.Round)
		return nil, err
	}

	setCurrent := false
	currentBib, err := getCurrentBib(l.ctx, l.svcCtx, &types.GetCurrentBibReq{Round: req.Round})
	fmt.Printf("CUrrent bib = %v, %v\n", currentBib, err)
	if err != nil || currentBib.CurrentBib == "" {
		setCurrent = true
		scl := SetBibCurrentLogic{
			Logger: l.Logger,
			ctx:    l.ctx,
			svcCtx: l.svcCtx,
		}
		scl.SetBibCurrent(&types.SetBibCurrentRequest{
			Round: req.Round,
			Bib:   req.Bib,
		})
	}

	bib := &types.Bib{
		ID:        req.Bib,
		Status:    "NONE",
		ISCurrent: false,
	}

	j, err := json.Marshal(bib)
	if err != nil {
		logx.Errorf("Unable to get json marshal %+v", bib)
		return nil, err
	}

	err = c.Write(l.ctx, []byte(req.Bib), j)
	if err != nil {
		logx.Errorf("Unable to Write %+v, err: %v", bib, err)
		return nil, err
	}

	if setCurrent {
		bib.ISCurrent = true
	}

	return bib, err
}
