package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachLogic {
	return &AttachLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachLogic) Attach(req *types.AttachRequest) (resp *types.AttachResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
