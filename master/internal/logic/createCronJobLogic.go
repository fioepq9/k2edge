package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCronJobLogic {
	return &CreateCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCronJobLogic) CreateCronJob(req *types.CreateCronJobRequest) (resp *types.CreateCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
