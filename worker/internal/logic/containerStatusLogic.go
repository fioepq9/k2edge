package logic

import (
	"context"
	"encoding/json"
	"io"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContainerStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerStatusLogic {
	return &ContainerStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainerStatusLogic) ContainerStatus(req *types.ContainerStatusRequest) (resp *types.ContainerStatusResponse, err error) {
	res, err := l.svcCtx.Docker.ContainerStats(l.ctx, req.ID, false)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	s, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resp = new(types.ContainerStatusResponse)
	data := make(map[string]interface{})
	json.Unmarshal(s, &data)
	resp.Status = data
	return resp, nil
}
