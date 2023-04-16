package logic

import (
	"context"
	"errors"
	"fmt"

	"k2edge/etcdutil"
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
	resp = new(types.ListJobResponse)

	if req.Namespace != "" {
		found, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Namespace)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, fmt.Errorf("job %s does not exist", req.Namespace)
		}
	}

	key := "/job"
	if req.Namespace != "" {
		key += "/" + req.Namespace 
	}

	jobs, err := etcdutil.GetOne[types.Job](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		if errors.Is(err, etcdutil.ErrKeyNotExist) {
			return resp, nil
		}
		return nil, err
	}

	for _, job := range *jobs {
		if req.Namespace == "" || job.Metadata.Namespace == req.Namespace {
			resp.Info = append(resp.Info, types.JobSimpleInfo{
				Namespace: job.Metadata.Namespace,
				Name: job.Metadata.Name,
				CreateTime: job.Config.CreateTime,
				Completions: job.Config.Completions,
				Succeeded: job.Succeeded,
				Schedule: job.Config.Schedule,
			})
		}
	}
	return resp, nil
}
