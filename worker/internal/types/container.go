package types

import "github.com/docker/docker/api/types"

type containerBuildOption func(*Container)

func NewContainer(opts ...containerBuildOption) Container {
	var c Container
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

type containerBuildOptions struct{}

var ContainerBuildOptions containerBuildOptions

func (containerBuildOptions) FromDocker(container types.Container) containerBuildOption {
	return func(c *Container) {
		// todo
		// c.config =
		// c.Status =
	}
}

func (containerBuildOptions) WithNamespace(namespace string) containerBuildOption {
	return func(c *Container) {
		c.Metadata.Namespace = namespace
	}
}
