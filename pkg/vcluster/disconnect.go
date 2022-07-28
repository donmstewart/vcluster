package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type DisconnectCommand struct {
	Name        string        `yaml:"name"`
	Namespace   string        `yaml:"namespace"`
	KubeContext string        `yaml:"kubeContext"`
	Arguments   []string      `yaml:"arguments,omitempty"`
	Flags       builder.Flags `yaml:"flags,omitempty"`
	Outputs     []Output      `yaml:"outputs,omitempty"`
}

func (c DisconnectCommand) GetSuffixArguments() []string {
	return nil
}

func (c DisconnectCommand) GetCommand() string {
	return "vcluster"
}

func (c DisconnectCommand) GetArguments() []string {
	// Final Command: docker pull carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	args := []string{
		"disconnect",
		c.Name,
		"--namespace",
		c.Namespace,
	}

	if c.KubeContext != "" {
		args = append(args, "--context")
		args = append(args, c.KubeContext)
	}

	args = append(args, c.Arguments...)

	return args
}

func (c DisconnectCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c DisconnectCommand) SuppressesOutput() bool {
	return false
}
