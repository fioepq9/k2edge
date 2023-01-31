package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CordonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCordonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CordonLogic {
	return &CordonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CordonLogic) Cordon(req *types.CordonRequest) (resp *types.CordonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
