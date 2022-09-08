package logic

import (
	"context"
	"fmt"

	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"

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

	uid := uuid.NewString()

	r := &types.Round{
		ID:     uid,
		Name:   req.Name,
		Status: "NONE",
	}

	err = l.svcCtx.Insert(r.ID, r)
	if err != nil {
		logx.Errorf("Error while Saving (%+v), err:%v", r, err)
		return nil, err
	}

	err = l.svcCtx.FindOne(r, bolthold.Where("ID").Eq(r.ID))
	// err = l.svcCtx.One("ID", r.ID, r)
	if err != nil {
		logx.Errorf("Strange, did not find the saved bib, err: %v", err)
		return nil, err
	}

	// current, err := GetCurrent(l.svcCtx.Store)
	// if err != nil || current.CurrentRound == "" {
	// 	err = SetCurrentRound(l.svcCtx.Store, r.ID)
	// 	if err != nil {
	// 		return nil,
	// 			err
	// 	}
	// 	r.ISCurrent = true
	// }

	return r, nil
}

func GetCurrent(db *bolthold.Store) (*types.Current, error) {
	current := new(types.Current)
	err := db.Get("current", current)
	if err != nil {
		return nil, err
	}
	fmt.Printf("XXXXXX current = %+v\n", current)
	return current, err
}

func SetCurrentRound(db *bolthold.Store, round string) error {
	current := new(types.Current)
	db.Get("current", current)
	current.CurrentRound = round
	return db.Update("current", current)
}

func SetCurrentStartBib(db *bolthold.Store, bib string) error {
	current := new(types.Current)
	db.Get("current", current)
	current.CurrentStartBib = bib
	return db.Update("current", current)

}

func SetCurrentEndBib(db *bolthold.Store, bib string) error {
	current := new(types.Current)
	db.Get("current", current)
	current.CurrentEndBib = bib
	return db.Update("current", current)
}
