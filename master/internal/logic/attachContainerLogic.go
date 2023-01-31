package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachContainerLogic {
	return &AttachContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachContainerLogic) AttachContainer(req *types.AttachContainerRequest) (resp *types.AttachContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
