// Code generated by goctl. DO NOT EDIT.
package types

type ClusterInfoResponse struct {
	Todo string `json:"todo"`
}

type CreateContainerRequest struct {
	Container Container `json:"container" yaml:"container"`
}

type CreateContainerResponse struct {
	ContainerInfo ContainerInfo `json:"containerInfo"`
}

type GetContainerRequest struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
}

type GetContainerResponse struct {
	Container Container `json:"container"`
}

type ListContainerRequest struct {
	Namespace string `form:"namespace,optional"`
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

type MigrateContainerRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Node      string `json:"node"`
}

type EventRequest struct {
	Message Message `json:"message"`
}

type Message struct {
	Status   string `json:"status,omitempty,optional"`
	ID       string `json:"id,omitempty,optional"`
	From     string `json:"from,omitempty,optional"`
	Type     string `json:"type"`
	Action   string `json:"action"` //create、start、die
	Actor    Actor  `json:"actor,optional"`
	Scope    string `json:"scope,omitempty,optional"`
	Time     int64  `json:"time,omitempty,optional"`
	TimeNano int64  `json:"timeNano,omitempty,optional"`
}

type Actor struct {
	ID         string      `json:"id,optional"`
	Attributes interface{} `json:"attributes,optional"`
}

type Metadata struct {
	Namespace string `json:"namespace" yaml:"namespace"`
	Kind      string `json:"kind" yaml:"kind"`
	Name      string `json:"name" yaml:"name"`
}

type Error struct {
	Todo string `json:"todo"`
}

type Container struct {
	Metadata        Metadata        `json:"metadata" yaml:"metadata"`
	ContainerConfig ContainerConfig `json:"container_config" yaml:"containerConfig"`
	ContainerStatus ContainerStatus `json:"container_status,optional" yaml:"containerStatus"`
}

type ContainerConfig struct {
	Deployment string           `json:"deployment,optional"`
	Job        string           `json:"job,optional"`
	Image      string           `json:"image" yaml:"image"`
	NodeName   string           `json:"node_name,optional" yaml:"nodeName"`
	Command    string           `json:"command,optional" yaml:"command"`
	Args       []string         `json:"args,optional" yaml:"args"`
	Expose     []ExposedPort    `json:"expose,optional" yaml:"expose"`
	Env        []string         `json:"env,optional" yaml:"env"`
	Limit      ContainerLimit   `json:"limit,optional" yaml:"limit"`
	Request    ContainerRequest `json:"request,optional" yaml:"request"`
}

type ContainerLimit struct {
	CPU    int64 `json:"cpu,default=50000000" yaml:"cpu"`
	Memory int64 `json:"memory,default=104857600" yaml:"memory"`
}

type ContainerRequest struct {
	CPU    int64 `json:"cpu,default=50000000" yaml:"cpu"`
	Memory int64 `json:"memory,default=104857600" yaml:"memory"`
}

type ExposedPort struct {
	Port     int64  `json:"port" yaml:"port"`
	Protocol string `json:"protocol" yaml:"protocol"`
	HostPort int64  `json:"host_port" yaml:"hostPort"`
}

type ContainerStatus struct {
	Status      string      `json:"status,optional" yaml:"status"`
	Node        string      `json:"node,optional" yaml:"node"`
	ContainerID string      `json:"container_id,optional" yaml:"containerID"`
	Info        interface{} `json:"info,optional" yaml:"info"`
}

type Job struct {
	Metadata  Metadata  `json:"metadata" yaml:"matadata"`
	Config    JobConfig `json:"config" yaml:"config"`
	Succeeded int       `json:"succeeded,optional" yaml:"succeeded"`
}

type JobConfig struct {
	CreateTime  int64             `json:"create_time,optional"`
	Completions int               `json:"completions" yaml:"completions"`
	Schedule    string            `json:"schedule,optional" yaml:"schedule"`
	Template    ContainerTemplate `json:"template" yaml:"template"`
}

type Deployment struct {
	Metadata Metadata         `json:"metadata" yaml:"metadata"`
	Config   DeploymentConfig `json:"config" yaml:"config"`
	Status   DeploymentStatus `json:"status,optional" yaml:"status"`
}

type DeploymentConfig struct {
	CreateTime int64             `json:"create_time,optional"`
	Replicas   int               `json:"replicas,default=1" yaml:"replicas"`
	Template   ContainerTemplate `json:"container_template" yaml:"containerTemplate"`
}

type ContainerTemplate struct {
	Name     string           `json:"name,optional" yaml:"name"`
	Image    string           `json:"image" yaml:"image"`
	NodeName string           `json:"node_name,optional" yaml:"nodeName"`
	Command  string           `json:"command,optional" yaml:"command"`
	Args     []string         `json:"args,optional" yaml:"args"`
	Expose   []ExposedPort    `json:"expose,optional" yaml:"expose"`
	Env      []string         `json:"env,optional" yaml:"env"`
	Limit    ContainerLimit   `json:"limit,optional" yaml:"limit"`
	Request  ContainerRequest `json:"request,optional" yaml:"request"`
}

type DeploymentStatus struct {
	AvailableReplicas int             `json:"available_replicas" yaml:"availableReplicas"`
	Containers        []ContainerInfo `json:"containers,optional" yaml:"containers"`
}

type ContainerInfo struct {
	Name        string `json:"name" yaml:"name"`
	Node        string `json:"node" yaml:"node"`
	ContainerID string `json:"containerID" yaml:"containerID"`
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
	WorkerURL string `json:"worker_url" yaml:"workerUrl"`
	MasterURL string `json:"master_url" yaml:"masterUrl"`
}

type Spec struct {
	Unschedulable bool     `json:"unschedulable"`
	Capacity      Capacity `json:"capacity" yaml:"capacity"`
}

type Status struct {
	Working     bool        `json:"working"`
	Allocatable Allocatable `json:"allocatable"`
	Condition   Condition   `json:"condition"`
}

type Capacity struct {
	CPU    int64 `json:"cpu" yaml:"cpu"`
	Memory int64 `json:"memory" yaml:"memory"`
}

type Allocatable struct {
	CPU    int64 `json:"cpu"`
	Memory int64 `json:"memory"`
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

type CreateDeploymentRequest struct {
	Deployment Deployment `json:"deployment" yaml:"deployment"`
}

type CreateDeploymentResponse struct {
	Err []string `json:"err"`
}

type GetDeploymentRequest struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
}

type GetDeploymentResponse struct {
	Deployment Deployment `json:"deployment"`
}

type ListDeploymentRequest struct {
	Namespace string `form:"namespace,optional"`
}

type ListDeploymentResponse struct {
	Info []DeploymentSimpleInfo `json:"info"`
}

type DeploymentSimpleInfo struct {
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Replicas   int    `json:"replicas"`
	Available  int    `json:"available"`
}

type DeleteDeploymentRequest struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	RemoveVolumnes bool   `json:"remove_volumns,optional"`
	RemoveLinks    bool   `json:"remoce_links,optional"`
	Force          bool   `json:"force,optional"`
	Timeout        int    `json:"timeout,optional"`
}

type DeleteDeploymentResponse struct {
	Err []string `json:"err"`
}

type ApplyDeploymentRequest struct {
	Namespace string           `json:"namespace" yaml:"namespace"`
	Name      string           `json:"name" yaml:"name"`
	Config    DeploymentConfig `json:"config" yaml:"config"`
}

type ApplyDeploymentResponse struct {
	Err []string `json:"err"`
}

type ScaleRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Replicas  int    `json:"replicas"`
}

type CreateJobRequest struct {
	Job Job `json:"job" yaml:"job"`
}

type GetJobRequest struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
}

type GetJobResponse struct {
	Job Job `json:"job"`
}

type ListJobRequest struct {
	Namespace string `form:"namespace,optional"`
}

type ListJobResponse struct {
	Info []JobSimpleInfo `json:"info"`
}

type JobSimpleInfo struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	CreateTime  int64  `json:"create_time"`
	Completions int    `json:"completions"`
	Succeeded   int    `json:"succeeded"`
	Schedule    string `json:"schedule"`
}

type DeleteJobRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type CreateNamespaceRequest struct {
	Name string `json:"name"`
}

type GetNamespaceRequest struct {
	Name string `form:"name"`
}

type GetNamespaceResponse struct {
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
	Name     string   `json:"name" yaml:"name"`
	Roles    []string `json:"roles" yaml:"roles"`
	BaseURL  NodeURL  `json:"base_url" yaml:"baseUrl"`
	Capacity Capacity `json:"capacity" yaml:"capacity"`
}

type NodeListRequest struct {
	All bool `form:"all, default=true"`
}

type NodeListResponse struct {
	NodeList []NodeList `json:"node_list"`
}

type NodeList struct {
	Name         string  `json:"name"`
	RegisterTime int64   `json:"register_time"`
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
