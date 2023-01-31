package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCronjobLogic {
	return &DeleteCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCronjobLogic) DeleteCronjob(req *types.DeleteCronjobRequest) (resp *types.DeleteCronjobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
