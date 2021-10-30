package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type DeleteCommand struct {
	Name            string        `yaml:"name"`
	Namespace       string        `yaml:"namespace"`
	KubeContext     string        `yaml:"kubeContext"`
	DeleteNamespace bool          `yaml:"deleteNamespace"`
	KeepPVC         bool          `yaml:"keepPVC"`
	Arguments       []string      `yaml:"arguments,omitempty"`
	Flags           builder.Flags `yaml:"flags,omitempty"`
	Outputs         []Output      `yaml:"outputs,omitempty"`
}

func (c DeleteCommand) GetSuffixArguments() []string {
	return nil
}

func (c DeleteCommand) GetCommand() string {
	return "vcluster"
}

func (c DeleteCommand) GetArguments() []string {
	// Final Command: docker pull carolynvs/zombies:v1.0 ARGUMENTS --FLAGS

	args := []string{
		"delete",
		c.Name,
		"--namespace",
		c.Namespace,
	}

	if c.KubeContext != "" {
		args = append(args, "--context")
		args = append(args, c.KubeContext)
	}

	if c.DeleteNamespace {
		args = append(args, "--delete-namespace")
	}

	if c.KeepPVC {
		args = append(args, "--keep-pvc")
	}

	args = append(args, c.Arguments...)

	return args
}

func (c DeleteCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c DeleteCommand) SuppressesOutput() bool {
	return false
}
