package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCronjobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCronjobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCronjobLogic {
	return &CreateCronjobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCronjobLogic) CreateCronjob(req *types.CreateCronjobRequest) (resp *types.CreateCronjobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
