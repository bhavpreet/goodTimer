package logic

import (
	"context"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
	"github.com/bhavpreet/goodTimer/parser"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"
	"github.com/timshannon/bolthold"

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

func (l *ListBibsLogic) ListBibs(req *types.ListBibsReq) (resp *types.ListBibsResp, err error) {
	// todo: add your logic here and delete this line

	var bibs []types.Bib
	err = l.svcCtx.Find(&bibs,
		bolthold.Where("ParentRoundID").Eq(req.Round).SortBy("EndTime").Reverse())
	if err != nil {
		logx.Errorf("Did not found round : %s", req.Round)
		return
	}

	current, _ := GetCurrent(l.svcCtx.Store)
	r, err := NewGetRoundLogic(l.ctx, l.svcCtx).GetRound(&types.GetRoundRequest{
		ID: req.Round,
	})
	if err != nil {
		return nil, err
	}

	for i := range bibs {
		if bibs[i].ID == current.CurrentStartBib {
			bibs[i].ISCurrentStart = true
		}
		if bibs[i].ID == current.CurrentEndBib {
			bibs[i].ISCurrentEnd = true
		}
		bibs[i].Round = *r

		var startTime time.Time
		var endTime time.Time
		var duration parser.Timespan
		// update Duration if start and finish are valid
		if bibs[i].StartTime != "" && bibs[i].StartTime != "DNS" {
			l := len(bibs[i].StartTime)
			startTime, err = time.Parse(timy2.ImpulseTimeFormat[:l], bibs[i].StartTime)
			if err != nil {
				logx.Errorf(
					"Unable to parse START time %s into format %s, err = %v",
					bibs[i].StartTime, timy2.ImpulseTimeFormat, err)
				continue // ignore
			}

			if bibs[i].EndTime != "" && bibs[i].EndTime != "DNF" {
				l := len(bibs[i].StartTime)
				endTime, err = time.Parse(timy2.ImpulseTimeFormat[:l], bibs[i].EndTime)
				if err != nil {
					logx.Errorf(
						"Unable to parse END time %s into format %s, err = %v",
						bibs[i].EndTime, timy2.ImpulseTimeFormat, err)
					continue // ignore
				}

				duration = parser.Timespan(endTime.Sub(startTime))
				bibs[i].Duration = duration.Format(parser.DurationFormat)
			}
		}
	}

	resp = new(types.ListBibsResp)
	if bibs == nil {
		bibs = []types.Bib{}
	}
	resp.Bibs = bibs
	return resp, nil
}
