package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UncordonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUncordonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UncordonLogic {
	return &UncordonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UncordonLogic) Uncordon(req *types.UncordonRequest) (resp *types.UncordonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
