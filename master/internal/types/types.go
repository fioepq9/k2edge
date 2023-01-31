// Code generated by goctl. DO NOT EDIT.
package types

type CreateCronjobRequest struct {
	Todo string `json:"todo"`
}

type CreateCronjobResponse struct {
	Error Error `json:"error,omitempty"`
}

type GetCronjobRequest struct {
	Todo string `json:"todo"`
}

type GetCronjobResponse struct {
	Response Metadata `json:"response"`
}

type DeleteCronjobRequest struct {
	Todo string `json:"todo"`
}

type DeleteCronjobResponse struct {
	Error Error `json:"error,omitempty"`
}

type ApplyCronjobRequest struct {
	Todo string `json:"todo"`
}

type ApplyCronjobResponse struct {
	Error Error `json:"error,omitempty"`
}

type ContainerConfig struct {
	Todo string `json:"todo"`
}

type ContainerStatus struct {
	Todo string `json:"todo"`
}

type ContainerMetadata struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type Container struct {
	Metadata ContainerMetadata `json:"metadata"`
	Config   ContainerConfig   `json:"config"`
	Status   ContainerStatus   `json:"status"`
}

type Command struct {
	Todo string `json:"todo"`
}

type Error struct {
	Todo string `json:"todo"`
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}
