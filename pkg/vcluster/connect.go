package vcluster

import (
	"strconv"

	"get.porter.sh/porter/pkg/exec/builder"
)

// var _ DockerCommand = PullCommand{}

type ConnectCommand struct {
	Name        string        `yaml:"name"`
	Namespace   string        `yaml:"namespace"`
	KubeConfig  string        `yaml:"kubeConfig"`
	KubeContext string        `yaml:"kubeContext"`
	Address     string        `yaml:"address"`
	LocalPort   int           `yaml:"localPort"`
	Pod         string        `yaml:"pod"`
	Server      string        `yaml:"server"`
	Arguments   []string      `yaml:"arguments,omitempty"`
	Flags       builder.Flags `yaml:"flags,omitempty"`
	Outputs     []Output      `yaml:"outputs,omitempty"`
}

func (c ConnectCommand) GetSuffixArguments() []string {
	return nil
}

func (c ConnectCommand) GetCommand() string {
	return "vcluster"
}

func (c ConnectCommand) GetArguments() []string {
	// Final Command: vcluster connect vcluster-1 --namespace host-vcluster-1 ARGUMENTS --FLAGS

	args := []string{
		"connect",
		c.Name,
		"--namespace",
		c.Namespace,
	}

	if c.KubeContext != "" {
		args = append(args, "--context")
		args = append(args, c.KubeContext)
	}

	if c.KubeConfig != "" {
		args = append(args, "--kube-config")
		args = append(args, c.KubeConfig)
	}

	if c.Address != "" {
		args = append(args, "--address")
		args = append(args, c.Address)
	}

	if c.LocalPort != 0 {
		args = append(args, "--local-port")
		args = append(args, strconv.Itoa(c.LocalPort))
	}

	if c.Pod != "" {
		args = append(args, "--pod")
		args = append(args, c.Pod)
	}

	if c.LocalPort != 0 {
		args = append(args, "--server")
		args = append(args, c.Server)
	}

	args = append(args, c.Arguments...)
	args = append(args, "&")

	return args
}

func (c ConnectCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c ConnectCommand) SuppressesOutput() bool {
	return false
}
