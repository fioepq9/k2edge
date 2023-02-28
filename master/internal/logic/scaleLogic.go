package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleLogic {
	return &ScaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScaleLogic) Scale(req *types.ScaleRequest) (resp *types.ScaleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
