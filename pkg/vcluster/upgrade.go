package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type UpgradeCommand struct {
	Name        string        `yaml:"name"`
	Namespace   string        `yaml:"namespace"`
	KubeContext string        `yaml:"kubeContext"`
	Version     string        `yaml:"version,omitempty"`
	Arguments   []string      `yaml:"arguments,omitempty"`
	Flags       builder.Flags `yaml:"flags,omitempty"`
	Outputs     []Output      `yaml:"outputs,omitempty"`
}

func (c UpgradeCommand) GetSuffixArguments() []string {
	return nil
}

func (c UpgradeCommand) GetCommand() string {
	return "vcluster"
}

func (c UpgradeCommand) GetArguments() []string {
	// Final Command: docker pull carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	args := []string{
		"upgrade",
		c.Name,
		"--namespace",
		c.Namespace,
	}

	if c.KubeContext != "" {
		args = append(args, "--context")
		args = append(args, c.KubeContext)
	}

	if c.Version != "" {
		args = append(args, "--version")
		args = append(args, c.Version)
	}

	args = append(args, c.Arguments...)

	return args
}

func (c UpgradeCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c UpgradeCommand) SuppressesOutput() bool {
	return false
}
