// Code generated by goctl. DO NOT EDIT.
package types

type VersionResponse struct {
	Version string `json:"version"`
}

type RunContainerRequest struct {
	Config ContainerConfig `json:"config"`
}

type RunContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type RemoveContainerRequest struct {
	Selector Metadata `json:"selector"`
}

type RemoveContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type StopContainerRequest struct {
	Selector Metadata `json:"selector"`
}

type StopContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type StartContainerRequest struct {
	Selector Metadata `json:"selector"`
}

type StartContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type ContainerStatusRequest struct {
	Selector Metadata `json:"selector"`
}

type ContainerStatusResponse struct {
	Container Container `json:"container"`
}

type ListContainersRequest struct {
	Todo string `json:"todo"`
}

type ListContainersResponse struct {
	Containers []Container `json:"containers"`
}

type ExecRequest struct {
	Selector Metadata `json:"selector"`
	Command  Command  `json:"command"`
}

type ExecResponse struct {
	Error Error `json:"error,omitempty"`
}

type AttachRequest struct {
	Selector Metadata `json:"selector"`
}

type AttachResponse struct {
	Error Error `json:"error,omitempty"`
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type Error struct {
	Todo string `json:"todo"`
}

type ContainerConfig struct {
	Todo string `json:"todo"`
}

type ContainerStatus struct {
	Todo string `json:"todo"`
}

type Container struct {
	Metadata Metadata        `json:"metadata"`
	Config   ContainerConfig `json:"config"`
	Status   ContainerStatus `json:"status"`
}

type CronJobConfig struct {
	Todo string `json:"todo"`
}

type CronJobStatus struct {
	Todo string `json:"todo"`
}

type CronJob struct {
	Metadata Metadata      `json:"metadata"`
	Config   CronJobConfig `json:"config"`
	Status   CronJobStatus `json:"status"`
}

type Command struct {
	Todo string `json:"todo"`
}

type Namespace struct {
	Name string `json:"name"`
}
