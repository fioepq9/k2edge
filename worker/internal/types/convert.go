package types

import (
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func (c ContainerConfig) DockerFormat() container.Config {
	exposedPorts := make(map[nat.Port]struct{})
	for _, k := range c.ExposedPorts {
		exposedPorts[nat.Port(k)] = struct{}{}
	}
	volumes := make(map[string]struct{})
	for _, k := range c.Volumes {
		volumes[k] = struct{}{}
	}
	var healthCheck *container.HealthConfig
	if c.Healthcheck != nil {
		healthCheck = &container.HealthConfig{
			Test:        c.Healthcheck.Test,
			Interval:    time.Duration(c.Healthcheck.Interval),
			Timeout:     time.Duration(c.Healthcheck.Timeout),
			StartPeriod: time.Duration(c.Healthcheck.StartPeriod),
			Retries:     c.Healthcheck.Retries,
		}
	}
	var labels map[string]string
	if l, ok := c.Labels.(map[string]string); ok {
		labels = l
	}
	return container.Config{
		Hostname:        c.Hostname,
		Domainname:      c.Domainname,
		User:            c.User,
		AttachStdin:     c.AttachStdin,
		AttachStdout:    c.AttachStdout,
		AttachStderr:    c.AttachStderr,
		ExposedPorts:    exposedPorts,
		Tty:             c.Tty,
		OpenStdin:       c.OpenStdin,
		StdinOnce:       c.StdinOnce,
		Env:             c.Env,
		Cmd:             c.Cmd,
		Healthcheck:     healthCheck,
		ArgsEscaped:     c.ArgsEscaped,
		Image:           c.Image,
		Volumes:         volumes,
		WorkingDir:      c.WorkingDir,
		Entrypoint:      c.Entrypoint,
		NetworkDisabled: c.NetworkDisabled,
		MacAddress:      c.MacAddress,
		OnBuild:         c.OnBuild,
		Labels:          labels,
		StopSignal:      c.StopSignal,
		StopTimeout:     &c.StopTimeout,
		Shell:           c.Shell,
	}
}
