// Code generated by goctl. DO NOT EDIT.
package types

type ClusterInfoResponse struct {
	Todo string `json:"todo"`
}

type CreateContainerRequest struct {
	Container Container `json:"container"`
}

type GetContainerRequest struct {
	Metadata Metadata `json:"metadata"`
}

type GetContainerResponse struct {
	Container Container `json:"container"`
}

type ListContainerResponse struct {
	ContainerSimpleInfo []ContainerSimpleInfo `json:"containers"`
}

type ContainerSimpleInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Node   string `json:"node"`
}

type DeleteContainerRequest struct {
	Metadata Metadata `json:"metadata"`
}

type ApplyContainerRequest struct {
	Todo string `json:"todo"`
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

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type Error struct {
	Todo string `json:"todo"`
}

type Container struct {
	Metadata        Metadata        `json:"metadata"`
	ContainerConfig ContainerConfig `json:"container_config"`
	ContainerStatus ContainerStatus `json:"container_status,optional"`
}

type ContainerConfig struct {
	Image         string        `json:"image"`
	NodeName      string        `json:"node_name,optional"`
	NodeNamespace string        `json:"node_namespace,optional"`
	Command       string        `json:"command,optional"`
	Args          []string      `json:"args,optional"`
	Expose        []ExposedPort `json:"expose,optional"`
	Env           []string      `json:"env,optional"`
}

type ExposedPort struct {
	Port     int64  `json:"port"`
	Protocol string `json:"protocol"`
	HostPort int64  `json:"host_port"`
}

type ContainerStatus struct {
	Status      string      `json:"status"`
	Node        string      `json:"node"`
	ContainerID string      `json:"container_id"`
	Info        interface{} `json:"info"`
}

type JobConfig struct {
	Todo string `json:"todo"`
}

type JobStatus struct {
	Todo string `json:"todo"`
}

type Job struct {
	Metadata              Metadata        `json:"metadata"`
	Node                  string          `json:"node"`
	Containers            []string        `json:"containers"`
	Completions           int64           `json:"completions"`
	BackoffLimit          int64           `json:"backoff_limit"`
	ActiveDeadlineSeconds int64           `json:"active_deadline_seconds"`
	StartTime             string          `json:"start_time"`
	CompletionTime        string          `json:"completion_time"`
	Active                int64           `json:"active"`
	Failed                int64           `json:"failed"`
	Succeeded             int64           `json:"succeeded"`
	Status                string          `json:"status"`
	Template              ContainerConfig `json:"template"`
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

type Node struct {
	Metadata     Metadata `json:"metadata"`
	Roles        []string `json:"roles"`
	BaseURL      NodeURL  `json:"base_url"`
	Status       string   `json:"status"`
	RegisterTime int64    `json:"register_time"`
}

type NodeURL struct {
	WorkerURL string `json:"worker_url"`
	MasterURL string `json:"master_url"`
}

type Command struct {
	Todo string `json:"todo"`
}

type Namespace struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	CreateTime int64  `json:"create_time"`
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
	Name string `json:"name"`
}

type GetNamespaceRequest struct {
	Name string `form:"name"`
}

type GetNamespaceResponse struct {
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    string `json:"age"`
}

type ListNamespaceRequest struct {
	All bool `form:"all, default=true"`
}

type ListNamespaceResponse struct {
	Namespaces []GetNamespaceResponse `json:"namespaces"`
}

type DeleteNamespaceRequest struct {
	Name string `json:"name"`
}

type RegisterRequest struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Roles     []string `json:"roles"`
	BaseURL   NodeURL  `json:"base_url"`
}

type NodeTopRequest struct {
	Name string `json:"name"`
}

type NodeTopResponse struct {
	TopInfo []TopInfo `json:"top_info"`
}

type TopInfo struct {
	Name       string `json:"name"`
	CPU        string `json:"CPU"`
	CPUProp    string `json:"CPU_prop"`
	Memory     string `json:"memory"`
	MemoryProp string `json:"memory_prop"`
}

type CordonRequest struct {
	Metadata Metadata `json:"metadata"`
}

type UncordonRequest struct {
	Metadata Metadata `json:"metadata"`
}

type DrainRequest struct {
	Metadata Metadata `json:"metadata"`
}

type DeleteRequest struct {
	Metadata Metadata `json:"metadata"`
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
