package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecLogic {
	return &ExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecLogic) Exec(req *types.ExecRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
