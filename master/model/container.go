package model

import "k2edge/master/internal/types"

type Container struct {
	types.Container
	ContainerStatus ContainerStatus `json:"container_status"`
}

type ContainerStatus struct {
	// 所在 node
	Node string `json:"node"`
	// 容器 ID
	ContainerID string `json:"container_id"`
	// 实际容器名
	ContainerName string `json:"container_name"`
	Status        interface{}
}
