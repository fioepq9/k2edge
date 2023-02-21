// Code generated by goctl. DO NOT EDIT.
package types

type ClusterInfoResponse struct {
	Todo string `json:"todo"`
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

type DeploymentConfig struct {
	Todo string `json:"todo"`
}

type DeploymentStatus struct {
	Todo string `json:"todo"`
}

type Deployment struct {
	Metadata Metadata         `json:"metadata"`
	Config   DeploymentConfig `json:"config"`
	Status   DeploymentStatus `json:"status"`
}

type JobConfig struct {
	Todo string `json:"todo"`
}

type JobStatus struct {
	Todo string `json:"todo"`
}

type Job struct {
	Metadata Metadata  `json:"metadata"`
	Config   JobConfig `json:"config"`
	Status   JobStatus `json:"status"`
}

type TokenConfig struct {
	Todo string `json:"todo"`
}

type TokenStatus struct {
	Todo string `json:"todo"`
}

type Token struct {
	Metadata Metadata    `json:"metadata"`
	Config   TokenConfig `json:"config"`
	Status   TokenStatus `json:"status"`
}

type Command struct {
	Todo string `json:"todo"`
}

type Namespace struct {
	Name string `json:"name"`
}

type CreateContainerRequest struct {
	Todo string `json:"todo"`
}

type CreateContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetContainerRequest struct {
	Todo string `json:"todo"`
}

type GetContainerResponse struct {
	Container Container `json:"container"`
}

type DeleteContainerRequest struct {
	Todo string `json:"todo"`
}

type DeleteContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type RunContainerRequest struct {
	Container Container `json:"container"`
}

type RunContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type ApplyContainerRequest struct {
	Todo string `json:"todo"`
}

type ApplyContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type HistoryContainerRequest struct {
	Todo string `json:"todo"`
}

type HistoryContainerResponse struct {
	Container Container `json:"container"`
}

type UndoContainerRequest struct {
	Todo string `json:"todo"`
}

type UndoContainerResponse struct {
	Error Error `json:"error,omitempty"`
}

type AttachContainerRequest struct {
	Todo string `json:"todo"`
}

type AttachContainerResponse struct {
	Todo string `json:"todo"`
}

type ExecContainerRequest struct {
	Todo string `json:"todo"`
}

type ExecContainerResponse struct {
	Todo string `json:"todo"`
}

type LogsContainerRequest struct {
	Todo string `json:"todo"`
}

type LogsContainerResponse struct {
	Todo string `json:"todo"`
}

type ContainerTopRequest struct {
	Selector Metadata `json:"selector"`
}

type ContainerTopResponse struct {
	Error Error `json:"error,omitempty"`
}

type CreateCronJobRequest struct {
	Todo string `json:"todo"`
}

type CreateCronJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetCronJobRequest struct {
	Todo string `json:"todo"`
}

type GetCronJobResponse struct {
	CronJob CronJob `json:"cronjob"`
}

type DeleteCronJobRequest struct {
	Todo string `json:"todo"`
}

type DeleteCronJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type ApplyCronJobRequest struct {
	Todo string `json:"todo"`
}

type ApplyCronJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type HistoryCronJobRequest struct {
	Todo string `json:"todo"`
}

type HistoryCronJobResponse struct {
	CronJob CronJob `json:"cronjob"`
}

type UndoCronJobRequest struct {
	Todo string `json:"todo"`
}

type UndoCronJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type LogsCronJobRequest struct {
	Todo string `json:"todo"`
}

type LogsCronJobResponse struct {
	Todo string `json:"todo"`
}

type CreateDeploymentRequest struct {
	Todo string `json:"todo"`
}

type CreateDeploymentResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetDeploymentRequest struct {
	Todo string `json:"todo"`
}

type GetDeploymentResponse struct {
	Deployment Deployment `json:"deployment"`
}

type DeleteDeploymentRequest struct {
	Todo string `json:"todo"`
}

type DeleteDeploymentResponse struct {
	Error Error `json:"error,omitempty"`
}

type ApplyDeploymentRequest struct {
	Todo string `json:"todo"`
}

type ApplyDeploymentResponse struct {
	Error Error `json:"error,omitempty"`
}

type HistoryDeploymentRequest struct {
	Todo string `json:"todo"`
}

type HistoryDeploymentResponse struct {
	Deployment Deployment `json:"deployment"`
}

type UndoDeploymentRequest struct {
	Todo string `json:"todo"`
}

type UndoDeploymentResponse struct {
	Error Error `json:"error,omitempty"`
}

type ScaleRequest struct {
	Todo string `json:"todo"`
}

type ScaleResponse struct {
	Todo string `json:"todo"`
}

type AttachDeploymentRequest struct {
	Todo string `json:"todo"`
}

type AttachDeploymentResponse struct {
	Todo string `json:"todo"`
}

type ExecDeploymentRequest struct {
	Todo string `json:"todo"`
}

type ExecDeploymentResponse struct {
	Todo string `json:"todo"`
}

type LogsDeploymentRequest struct {
	Todo string `json:"todo"`
}

type LogsDeploymentResponse struct {
	Todo string `json:"todo"`
}

type CreateJobRequest struct {
	Todo string `json:"todo"`
}

type CreateJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetJobRequest struct {
	Todo string `json:"todo"`
}

type GetJobResponse struct {
	Job Job `json:"job"`
}

type DeleteJobRequest struct {
	Todo string `json:"todo"`
}

type DeleteJobResponse struct {
	Error Error `json:"error,omitempty"`
}

type LogsJobRequest struct {
	Todo string `json:"todo"`
}

type LogsJobResponse struct {
	Todo string `json:"todo"`
}

type CreateNamespaceRequest struct {
	Todo string `json:"todo"`
}

type CreateNamespaceResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetNamespaceRequest struct {
	Todo string `json:"todo"`
}

type GetNamespaceResponse struct {
	Namespace Namespace `json:"namespace"`
}

type ListNamespacesRequest struct {
	Todo string `json:"todo"`
}

type ListNamespacesResponse struct {
	Namespaces []Namespace `json:"namespaces"`
}

type DeleteNamespaceRequest struct {
	Todo string `json:"todo"`
}

type DeleteNamespaceResponse struct {
	Error Error `json:"error,omitempty"`
}

type NodeTopRequest struct {
	Selector Metadata `json:"selector"`
}

type NodeTopResponse struct {
	Error Error `json:"error,omitempty"`
}

type CordonRequest struct {
	Selector Metadata `json:"selector"`
}

type CordonResponse struct {
	Error Error `json:"error,omitempty"`
}

type UncordonRequest struct {
	Selector Metadata `json:"selector"`
}

type UncordonResponse struct {
	Error Error `json:"error,omitempty"`
}

type DrainRequest struct {
	Selector Metadata `json:"selector"`
}

type DrainResponse struct {
	Error Error `json:"error,omitempty"`
}

type CreateTokenRequest struct {
	Todo string `json:"todo"`
}

type CreateTokenResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetTokenRequest struct {
	Todo string `json:"todo"`
}

type GetTokenResponse struct {
	Token Token `json:"token"`
}

type DeleteTokenRequest struct {
	Todo string `json:"todo"`
}

type DeleteTokenResponse struct {
	Error Error `json:"error,omitempty"`
}

type ApplyTokenRequest struct {
	Todo string `json:"todo"`
}

type ApplyTokenResponse struct {
	Error Error `json:"error,omitempty"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
