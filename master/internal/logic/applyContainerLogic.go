package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyContainerLogic {
	return &ApplyContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyContainerLogic) ApplyContainer(req *types.ApplyContainerRequest) (resp *types.ApplyContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
