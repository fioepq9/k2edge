// Code generated by goctl. DO NOT EDIT.
package client

type VersionResponse struct {
	Version string `json:"version"`
}

type CreateContainerRequest struct {
	ContainerName string          `json:"container_name"`
	Config        ContainerConfig `json:"config"`
}

type CreateContainerResponse struct {
	ID string `json:"id"`
}

type RemoveContainerRequest struct {
	ID            string `json:"id"`
	RemoveVolumes bool   `json:"remove_volumnes,optional"`
	RemoveLinks   bool   `json:"remove_links,optional"`
	Force         bool   `json:"force"`
}

type StopContainerRequest struct {
	ID      string `json:"id"`
	Timeout int    `json:"timeout,optional"`
}

type StartContainerRequest struct {
	ID            string `json:"id"`
	CheckpointID  string `json:"checkpoint_id,optional"`
	CheckpointDir string `json:"checkpoint_dir,optional"`
}

type ContainerStatusRequest struct {
	ID string `form:"id"`
}

type ContainerStatusResponse struct {
	Status interface{} `json:"status"`
}

type ListContainersRequest struct {
	Size   bool   `form:"size,optional"`
	All    bool   `form:"all,optional"`
	Latest bool   `form:"latest,optional"`
	Since  string `form:"since,optional"`
	Before string `form:"before,optional"`
	Limit  int    `form:"limit,optional"`
}

type ListContainersResponse struct {
	Containers interface{} `json:"containers"`
}

type ExecRequest struct {
	Container    string   `form:"container"`
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

type AttachRequest struct {
	Container  string `form:"container"`
	Stream     bool   `form:"stream,default=true"`
	Stdin      bool   `form:"stdin,default=true"`
	Stdout     bool   `form:"stdout,default=true"`
	Stderr     bool   `form:"stderr,default=true"`
	DetachKeys string `form:"detach_keys,optional"`
	Logs       bool   `form:"logs,optional"`
}

type NodeTopResponse struct {
	Images            []string  `json:"images"`
	CPU               []CPUInfo `json:"cpu"`
	MemoryUsed        uint64    `json:"memory_used"`
	MemoryAvailable   uint64    `json:"memory_available"`
	MemoryUsedPercent float64   `json:"memory_used_percent"`
	MemoryTotal       uint64    `json:"memory_total"`
	DiskUsed          uint64    `json:"disk_used"`
	DiskFree          uint64    `json:"disk_free"`
	DiskUsedPercent   float64   `json:"disk_used_percent"`
	DiskTotal         uint64    `json:"disk_total"`
}

type CPUInfo struct {
	CPU       int32   `json:"cpu"`
	Cores     int32   `json:"cores"`
	Mhz       float64 `json:"mhz"`
	ModelName string  `json:"model_name"`
	Percent   float64 `json:"percent"`
}

type LogsRequest struct {
	Container  string `form:"container"`
	ShowStdout bool   `form:"show_stdout,optional"`
	ShowStderr bool   `form:"show_stderr,optional"`
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
	Image    string        `json:"image"`
	NodeName string        `json:"node_name,optional"`
	Command  string        `json:"command,optional"`
	Args     []string      `json:"args,optional"`
	Expose   []ExposedPort `json:"expose,optional"`
	Env      []string      `json:"env,optional"`
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

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
