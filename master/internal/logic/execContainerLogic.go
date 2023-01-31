package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecContainerLogic {
	return &ExecContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecContainerLogic) ExecContainer(req *types.ExecContainerRequest) (resp *types.ExecContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
