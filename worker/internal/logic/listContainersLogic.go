package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContainersLogic {
	return &ListContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContainersLogic) ListContainers(req *types.ListContainersRequest) (resp *types.ListContainersResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
