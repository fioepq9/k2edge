// Code generated by goctl. DO NOT EDIT.
package types

type VersionResponse struct {
	Version string `json:"version"`
}

type RunContainerRequest struct {
	ContainerName string          `json:"container_name"`
	Config        ContainerConfig `json:"config"`
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
	Container string     `json:"container"`
	Config    ExecConfig `json:"config,optional"`
}

type ExecConfig struct {
	User         string   `json:"user"`          // User that will run the command
	Privileged   bool     `json:"privileged"`    // Is the container in privileged mode
	Tty          bool     `json:"tty"`           // Attach standard streams to a tty.
	AttachStdin  bool     `json:"attach_stdin"`  // Attach the standard input, makes possible user interaction
	AttachStderr bool     `json:"attach_stderr"` // Attach the standard error
	AttachStdout bool     `json:"attach_stdout"` // Attach the standard output
	Detach       bool     `json:"detach"`        // Execute in detach mode
	DetachKeys   string   `json:"detach_keys"`   // Escape keys for detach
	Env          []string `json:"env"`           // Environment variables
	WorkingDir   string   `json:"working_dir"`   // Working directory
	Cmd          []string `json:"cmd"`           // Execution commands and args
}

type AttachRequest struct {
	Container string       `json:"container"`
	Config    AttachConfig `json:"config,optional"`
}

type AttachConfig struct {
	Stream     bool   `json:"stream"`
	Stdin      bool   `json:"stdin"`
	Stdout     bool   `json:"stdout"`
	Stderr     bool   `json:"stderr"`
	DetachKeys string `json:"detach_keys"`
	Logs       bool   `json:"logs"`
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type Error struct {
	Todo string `json:"todo"`
}

type HealthConfig struct {
	Test        []string `json:"test"`
	Interval    int64    `json:"interval"`              // Interval is the time to wait between checks.
	Timeout     int64    `json:"timeout"`               // Timeout is the time to wait before considering the check to have hung.
	StartPeriod int64    `json:"start_period,optional"` // The start period for the container to initialize before the retries starts to count down.
	Retries     int      `json:"retries,optional"`
}

type ContainerConfig struct {
	Hostname        string        `json:"hostname,optional"`         // Hostname
	Domainname      string        `json:"domainname,optional"`       // Domainname
	User            string        `json:"user,optional"`             // User that will run the command(s) inside the container, also support user:group
	AttachStdin     bool          `json:"attach_stdin,optional"`     // Attach the standard input, makes possible user interaction
	AttachStdout    bool          `json:"attach_stdout,optional"`    // Attach the standard output
	AttachStderr    bool          `json:"attach_stderr,optional"`    // Attach the standard error
	ExposedPorts    []string      `json:"exposed_ports,optional"`    // List of exposed ports
	Tty             bool          `json:"tty,optional"`              // Attach standard streams to a tty, including stdin if it is not closed.
	OpenStdin       bool          `json:"open_stdin,optional"`       // Open stdin
	StdinOnce       bool          `json:"stdin_once,optional"`       // If true, close stdin after the 1 attached client disconnects.
	Env             []string      `json:"env,optional"`              // List of environment variable to set in the container
	Cmd             []string      `json:"cmd,optional"`              // Command to run when starting the container
	Healthcheck     *HealthConfig `json:"healthcheck,optional"`      // Healthcheck describes how to check the container is healthy
	ArgsEscaped     bool          `json:"args_escaped,optional"`     // True if command is already escaped (meaning treat as a command line) (Windows specific).
	Image           string        `json:"image"`                     // Name of the image as it was passed by the operator (e.g. could be symbolic)
	Volumes         []string      `json:"volumes,optional"`          // List of volumes (mounts) used for the container
	WorkingDir      string        `json:"working_dir,optional"`      // Current directory (PWD) in the command will be launched
	Entrypoint      []string      `json:"entrypoint,optional"`       // Entrypoint to run when starting the container
	NetworkDisabled bool          `json:"network_disabled,optional"` // Is network disabled
	MacAddress      string        `json:"mac_address,optional"`      // Mac Address of the container
	OnBuild         []string      `json:"on_build,optional"`         // ONBUILD metadata that were defined on the image Dockerfile
	Labels          interface{}   `json:"labels,optional"`           // List of labels set to this container
	StopSignal      string        `json:"stop_signal,optional"`      // Signal to stop a container
	StopTimeout     int           `json:"stop_timeout,optional"`     // Timeout (in seconds) to stop a container
	Shell           []string      `json:"shell,optional"`            // Shell for shell-form of RUN, CMD, ENTRYPOINT
}

type ContainerStatus struct {
	Todo string `json:"todo"`
}

type Container struct {
	Metadata          Metadata          `json:"metadata"`
	ContainerTemplate ContainerTemplate `json:"container_template"`
}

type ContainerTemplate struct {
	Image   string        `json:"image"`
	Command string        `json:"command"`
	Args    []string      `json:"args"`
	Expose  []ExposedPort `json:"expose"`
	Env     []string      `json:"env"`
}

type ExposedPort struct {
	Port     int64  `json:"port"`
	Protocol string `json:"protocol"`
	HostPort int64  `json:"host_port"`
}

type JobConfig struct {
	Todo string `json:"todo"`
}

type JobStatus struct {
	Todo string `json:"todo"`
}

type Job struct {
	Metadata              Metadata          `json:"metadata"`
	Node                  string            `json:"node"`
	Containers            []string          `json:"containers"`
	Completions           int64             `json:"completions"`
	BackoffLimit          int64             `json:"backoff_limit"`
	ActiveDeadlineSeconds int64             `json:"active_deadline_seconds"`
	StartTime             string            `json:"start_time"`
	CompletionTime        string            `json:"completion_time"`
	Active                int64             `json:"active"`
	Failed                int64             `json:"failed"`
	Succeeded             int64             `json:"succeeded"`
	Status                string            `json:"status"`
	Template              ContainerTemplate `json:"template"`
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
	BaseURL      string   `json:"base_url"`
	Status       string   `json:"status"`
	RegisterTime int64    `json:"register_time"`
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
