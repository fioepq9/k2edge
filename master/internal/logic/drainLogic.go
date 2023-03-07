package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrainLogic {
	return &DrainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrainLogic) Drain(req *types.DrainRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
