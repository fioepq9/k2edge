package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyTokenLogic {
	return &ApplyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyTokenLogic) ApplyToken(req *types.DeleteTokenRequest) (resp *types.DeleteTokenResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
