package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyCronjobLogic {
	return &ApplyCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyCronjobLogic) ApplyCronjob(req *types.DeleteCronjobRequest) (resp *types.DeleteCronjobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
