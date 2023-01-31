package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContainerLogic {
	return &GetContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContainerLogic) GetContainer(req *types.GetContainerRequest) (resp *types.GetContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
