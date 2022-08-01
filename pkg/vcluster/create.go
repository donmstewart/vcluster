package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

type CreateCommand struct {
	Name               string        `yaml:"name"`
	Namespace          string        `yaml:"namespace"`
	KubeContext        string        `yaml:"kubeContext"`
	CreateClusterRole  bool          `yaml:"createClusterRole"`
	CreateNamespace    bool          `yaml:"createNamespace"`             // If true the namespace will be created if it does not exist (default true)
	DisableIngressSync bool          `yaml:"disableIngressSync"`          // If true the virtual cluster will not sync any ingresses
	Upgrade            bool          `yaml:"upgrade"`                     // If true will try to upgrade the vcluster instead of failing if it already exists
	Distro             string        `yaml:"distro,omitempty"`            // Kubernetes distro to use for the virtual cluster. Allowed distros: k3s, k0s, k8s, eks (default "k3s")
	KubernetesVersion  string        `yaml:"kubernetesVersion,omitempty"` // The kubernetes version to use (e.g. v1.20). Patch versions are not supported
	Connect            bool          `yaml:"connect"`
	Expose             bool          `yaml:"expose"`
	Isolate            bool          `yaml:"isolate"`
	ChartName          string        `yaml:"chartName,omitempty"`    // The virtual cluster chart name to use (default "vcluster")
	ChartRepo          string        `yaml:"chartRepo,omitempty"`    // The virtual cluster chart repo to use (default "https://charts.loft.sh")
	ChartVersion       string        `yaml:"chartVersion,omitempty"` // The virtual cluster chart version to use (e.g. v0.9.1) (default "0.10.2")
	ExtraConfig        string        `yaml:"extraConfig"`
	Arguments          []string      `yaml:"arguments,omitempty"`
	Flags              builder.Flags `yaml:"flags,omitempty"`
	Outputs            []Output      `yaml:"outputs,omitempty"`
	// DEPRECATED: use --extra-values instead
	// K3SImage           string        `yaml:"k3sImage"`
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

	if c.ChartName != "" {
		args = append(args, "--chart-name")
		args = append(args, c.ChartName)
	}

	if c.ChartRepo != "" {
		args = append(args, "--chart-repo")
		args = append(args, c.ChartRepo)
	}

	if c.ChartVersion != "" {
		args = append(args, "--chart-version")
		args = append(args, c.ChartVersion)
	}

	// DEPRECATED: use --extra-values instead
	// if c.K3SImage != "" {
	// 	args = append(args, "--k3s-image string")
	// 	args = append(args, c.K3SImage)
	// }

	if c.Distro != "" {
		args = append(args, "--distro")
		args = append(args, c.Distro)
	}

	if c.KubernetesVersion != "" {
		args = append(args, "--kubernetes-version")
		args = append(args, c.Distro)
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

	if c.Isolate {
		args = append(args, "--isolate")
	}

	if c.Upgrade {
		args = append(args, "--upgrade")
	}

	if c.Expose {
		args = append(args, "--expose")
	}

	if c.ExtraConfig != "" {
		args = append(args, "-f")
		args = append(args, c.ExtraConfig)
	}

	args = append(args, c.Arguments...)

	// if c.Connect {
	// 	args = append(args, "&")
	// }

	return args
}

func (c CreateCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c CreateCommand) SuppressesOutput() bool {
	return false
}
