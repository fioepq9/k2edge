package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListJobLogic {
	return &ListJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListJobLogic) ListJob(req *types.ListJobRequest) (resp *types.ListJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
