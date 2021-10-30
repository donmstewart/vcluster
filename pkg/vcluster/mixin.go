//go:generate packr2
package vcluster

import (
	"get.porter.sh/porter/pkg/context"
)

const (
	defaultVClusterClientVersion string = "v0.4.3"
	defaultClientPlatform        string = "linux"
	defaultClientArchitecture    string = "amd64"
	defaultHelmClientVersion     string = "v3.3.4"
	defaultKubectlClientVersion  string = "v1.22.1"
	// clientVersionConstraint represents the semver constraint for the Helm client version
	// Currently, this mixin only supports Helm clients versioned v3.x.x
	helmClientVersionConstraint string = "^v3.x"
)

type Mixin struct {
	*context.Context
	ClientVersion        string `yaml:"clientVersion,omitempty"`
	HelmClientVersion    string `yaml:"helmClientVersion,omitempty"`
	KubectlClientVersion string `yaml:"kubectlClientVersion,omitempty"`
	ClientPlatform       string `yaml:"clientPlatform,omitempty"`
	ClientArchitecture   string `yaml:"clientArchitecture,omitempty"`
	// add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context:              context.New(),
		ClientVersion:        defaultVClusterClientVersion,
		HelmClientVersion:    defaultHelmClientVersion,
		ClientPlatform:       defaultClientPlatform,
		ClientArchitecture:   defaultClientArchitecture,
		KubectlClientVersion: defaultKubectlClientVersion,
	}, nil

}
