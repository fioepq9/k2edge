package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJobLogic {
	return &GetJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJobLogic) GetJob(req *types.GetJobRequest) (resp *types.GetJobResponse, err error) {
	key := etcdutil.GenerateKey("job", req.Namespace, req.Name)
	// 判断 job 是否存在, 存在则获取 job 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("job %s does not exist", req.Name)
	}

	//获取job信息
	job, err := etcdutil.GetOne[types.Job](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	resp = new(types.GetJobResponse)
	resp.Job = (*job)[0]

	return resp, nil
}
