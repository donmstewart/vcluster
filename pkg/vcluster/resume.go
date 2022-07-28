package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type ResumeCommand struct {
	Name        string        `yaml:"name"`
	Namespace   string        `yaml:"namespace"`
	KubeContext string        `yaml:"kubeContext"`
	Arguments   []string      `yaml:"arguments,omitempty"`
	Flags       builder.Flags `yaml:"flags,omitempty"`
	Outputs     []Output      `yaml:"outputs,omitempty"`
}

func (c ResumeCommand) GetSuffixArguments() []string {
	return nil
}

func (c ResumeCommand) GetCommand() string {
	return "vcluster"
}

func (c ResumeCommand) GetArguments() []string {
	// Final Command: docker pull carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	args := []string{
		"resume",
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

func (c ResumeCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c ResumeCommand) SuppressesOutput() bool {
	return false
}
