package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type CreateCommand struct {
	Name               string        `yaml:"name"`
	Namespace          string        `yaml:"namespace"`
	KubeContext        string        `yaml:"kubeContext"`
	CreateClusterRole  bool          `yaml:"createClusterRole"`
	CreateNamespace    bool          `yaml:"createNamespace"`
	DisableIngressSync bool          `yaml:"disableIngressSync"`
	Upgrade            bool          `yaml:"upgrade"`
	K3SImage           string        `yaml:"k3sImage"`
	Connect            bool          `yaml:"connect"`
	Arguments          []string      `yaml:"arguments,omitempty"`
	Flags              builder.Flags `yaml:"flags,omitempty"`
	Outputs            []Output      `yaml:"outputs,omitempty"`
}

func (c CreateCommand) GetSuffixArguments() []string {
	return nil
}

func (c CreateCommand) GetCommand() string {
	return "vcluster"
}

func (c CreateCommand) GetArguments() []string {
	// Final Command: vcluster vcluster-1 --namespace host-vcluster-1  ARGUMENTS --FLAGS

	args := []string{
		"create",
		c.Name,
		"--namespace",
		c.Namespace,
	}

	if c.KubeContext != "" {
		args = append(args, "--context")
		args = append(args, c.KubeContext)
	}

	if c.CreateClusterRole {
		args = append(args, "--create-cluster-role")
	}

	if c.K3SImage != "" {
		args = append(args, "--k3s-image string")
		args = append(args, c.K3SImage)
	}

	if c.Connect {
		args = append(args, "--connect")
	}

	if c.CreateNamespace {
		args = append(args, "--create-namespace")
	}

	if c.DisableIngressSync {
		args = append(args, "--disable-ingress-sync")
	}

	if c.Upgrade {
		args = append(args, "--upgrade")
	}

	args = append(args, c.Arguments...)

	if c.Connect {
		args = append(args, "&")
	}

	return args
}

func (c CreateCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c CreateCommand) SuppressesOutput() bool {
	return false
}
