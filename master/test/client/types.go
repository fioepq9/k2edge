// Code generated by goctl. DO NOT EDIT.
package client

type ClusterInfoResponse struct {
	Todo string `json:"todo"`
}

type CreateContainerRequest struct {
	Container Container `json:"container"`
}

type GetContainerRequest struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
}

type GetContainerResponse struct {
	Container Container `json:"container"`
}

type ListContainerRequest struct {
	Namespace string `form:"namespace"`
}

type ListContainerResponse struct {
	ContainerSimpleInfo []ContainerSimpleInfo `json:"containers"`
}

type ContainerSimpleInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
	Node      string `json:"node"`
}

type DeleteContainerRequest struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	RemoveVolumnes bool   `json:"remove_volumns,optional"`
	RemoveLinks    bool   `json:"remoce_links,optional"`
	Force          bool   `json:"force,optional"`
	Timeout        int    `json:"timeout,optional"`
}

type ApplyContainerRequest struct {
	Container Container `json:"container"`
}

type AttachContainerRequest struct {
	Namespace  string `form:"namespace"`
	Name       string `form:"name"`
	Stream     bool   `form:"stream,optional"`
	Stdin      bool   `form:"stdin,optional"`
	Stdout     bool   `form:"stdout,optional"`
	Stderr     bool   `form:"stderr,optional"`
	DetachKeys string `form:"detach_keys,optional"`
	Logs       bool   `form:"logs,optional"`
}

type ExecContainerRequest struct {
	Namespace    string   `form:"namespace"`
	Name         string   `form:"name"`
	User         string   `form:"user,optional"`          // User that will run the command
	Privileged   bool     `form:"privileged,optional"`    // Is the container in privileged mode
	Tty          bool     `form:"tty,optional"`           // Attach standard streams to a tty.
	AttachStdin  bool     `form:"attach_stdin,optional"`  // Attach the standard input, makes possible user interaction
	AttachStderr bool     `form:"attach_stderr,optional"` // Attach the standard error
	AttachStdout bool     `form:"attach_stdout,optional"` // Attach the standard output
	Detach       bool     `form:"detach,optional"`        // Execute in detach mode
	DetachKeys   string   `form:"detach_keys,optional"`   // Escape keys for detach
	Env          []string `form:"env,optional"`           // Environment variables
	WorkingDir   string   `form:"working_dir,optional"`   // Working directory
	Cmd          []string `form:"cmd"`                    // Execution commands and args
}

type LogsContainerRequest struct {
	Namespace  string `form:"namespace"`
	Name       string `form:"name"`
	Since      string `form:"since,optional"`
	Until      string `form:"until,optional"`
	Timestamps bool   `form:"timestamps,optional"`
	Follow     bool   `form:"follow,optional"`
	Tail       string `form:"tail,optional"`
	Details    bool   `form:"details,optional"`
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
	Image    string           `json:"image"`
	NodeName string           `json:"node_name,optional"`
	Command  string           `json:"command,optional"`
	Args     []string         `json:"args,optional"`
	Expose   []ExposedPort    `json:"expose,optional"`
	Env      []string         `json:"env,optional"`
	Limit    ContainerLimit   `json:"limit,optional"`
	Request  ContainerRequest `json:"request,optional"`
}

type ContainerLimit struct {
	CPU    int64 `json:"cpu,default=50000000"`
	Memory int64 `json:"memory,default=104857600"`
}

type ContainerRequest struct {
	CPU    int64 `json:"cpu,default=50000000"`
	Memory int64 `json:"memory,default=104857600"`
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
	Spec         Spec     `json:"spec"`
	RegisterTime int64    `json:"register_time"`
	Status       Status   `json:"status"`
}

type NodeURL struct {
	WorkerURL string `json:"worker_url"`
	MasterURL string `json:"master_url"`
}

type Spec struct {
	Unschedulable bool `json:"unschedulable"`
}

type Status struct {
	Working     bool        `json:"working"`
	Capacity    Capacity    `json:"capacity"`
	Allocatable Allocatable `json:"allocatable"`
	Condition   Condition   `json:"condition"`
}

type Capacity struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type Allocatable struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type Condition struct {
	Ready ConditionInfo `json:"ready"`
}

type ConditionInfo struct {
	Status            bool   `json:"status"`
	LastHeartbeatTime string `json:"LastHeartbeatTime"`
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

type ScaleRequest struct {
	Todo string `json:"todo"`
}

type ScaleResponse struct {
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
	Name    string   `json:"name"`
	Roles   []string `json:"roles"`
	BaseURL NodeURL  `json:"base_url"`
}

type NodeListRequest struct {
	All bool `form:"all, default=true"`
}

type NodeListResponse struct {
	NodeList []NodeList `json:"node_list"`
}

type NodeList struct {
	Name         string  `json:"name"`
	RegisterTime int64   `json:"register_name"`
	Status       string  `json:"status"`
	Roles        string  `json:"roles"`
	URL          NodeURL `json:"url"`
}

type NodeTopRequest struct {
	Name string `form:"name"`
}

type NodeTopResponse struct {
	Images            []string `json:"images"`
	CPUUsed           float64  `json:"cpu_used"`
	CPUFree           float64  `json:"cpu_free"`
	CPUTotal          float64  `json:"cpu_total"`
	CPUUsedPercent    float64  `json:"cpu_used_percent"`
	MemoryUsed        uint64   `json:"memory_used"`
	MemoryAvailable   uint64   `json:"memory_available"`
	MemoryUsedPercent float64  `json:"memory_used_percent"`
	MemoryTotal       uint64   `json:"memory_total"`
	DiskUsed          uint64   `json:"disk_used"`
	DiskFree          uint64   `json:"disk_free"`
	DiskUsedPercent   float64  `json:"disk_used_percent"`
	DiskTotal         uint64   `json:"disk_total"`
}

type CordonRequest struct {
	Name string `json:"name"`
}

type UncordonRequest struct {
	Name string `json:"name"`
}

type DrainRequest struct {
	Name string `json:"name"`
}

type DeleteRequest struct {
	Name string `json:"name"`
}

type ScheduleRequest struct {
	Name  string  `json:"name"`
	Ports []int64 `json:"posts"`
}

type ScheduleResponse struct {
	Images            []string `json:"images"`
	MemoryUsed        uint64   `json:"memory_used"`
	MemoryAvailable   uint64   `json:"memory_available"`
	MemoryUsedPercent float64  `json:"memory_used_percent"`
	MemoryTotal       uint64   `json:"memory_total"`
	PortUsable        bool     `json:"port_usable"`
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
