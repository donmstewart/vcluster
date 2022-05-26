package vcluster

import (
	"bytes"
	"fmt"
	"text/template"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	vClusterDockerfileLines string = `RUN curl -L --output vcluster https://github.com/loft-sh/vcluster/releases/download/{{ .ClientVersion }}/vcluster-{{ .ClientPlatform }}-{{ .ClientArchitecture }} && \
mv ./vcluster /usr/local/bin/vcluster && \
chmod a+x /usr/local/bin/vcluster
`
	kubectlDockerfileLines string = `RUN curl -o kubectl https://storage.googleapis.com/kubernetes-release/release/{{ .KubectlClientVersion }}/bin/{{ .ClientPlatform }}/{{ .ClientArchitecture }}/kubectl && \
mv kubectl /usr/local/bin && chmod a+x /usr/local/bin/kubectl
`
	helmDockerfileLines string = `ENV HELM_EXPERIMENTAL_OCI=1
RUN curl -L https://get.helm.sh/helm-{{ .HelmClientVersion }}-{{ .ClientPlatform }}-{{ .ClientArchitecture }}.tar.gz --output helm3.tar.gz && \
tar -xvf helm3.tar.gz && rm helm3.tar.gz && mv {{ .ClientPlatform }}-{{ .ClientArchitecture }}/helm /usr/local/bin/helm
`
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the vcluster mixin in porter.yaml
// mixins:
// - vcluster:
//	  clientVersion: "v0.0.0"

type MixinConfig struct {
	ClientVersion        string `yaml:"clientVersion,omitempty"`
	HelmClientVersion    string `yaml:"helmClientVersion,omitempty"`
	ClientPlatform       string `yaml:"clientPlatform,omitempty"`
	ClientArchitecture   string `yaml:"clientArchitecture,omitempty"`
	KubectlClientVersion string `yaml:"kubectlClientVersion,omitempty"`
}

// This is an example. Replace the following with whatever steps are needed to
// install required components into
// const vClusterDockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	// Create new Builder
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	vclusterTemplate, err := template.New("vclusterCmd").Parse(vClusterDockerfileLines)
	if err != nil {
		return err
	}

	kubectlTemplate, err := template.New("kubectlCmd").Parse(kubectlDockerfileLines)
	if err != nil {
		return err
	}

	helmTemplate, err := template.New("kubectlCmd").Parse(helmDockerfileLines)
	if err != nil {
		return err
	}

	helmSuppliedClientVersion := input.Config.HelmClientVersion
	if helmSuppliedClientVersion != "" {
		ok, err := validate(helmSuppliedClientVersion, helmClientVersionConstraint)
		if err != nil {
			return err
		}
		if !ok {
			return errors.Errorf("supplied clientVersion %q does not meet semver constraint %q",
				helmSuppliedClientVersion, helmClientVersionConstraint)
		}
		m.HelmClientVersion = helmSuppliedClientVersion
	}

	if input.Config.ClientVersion != "" {
		m.ClientVersion = input.Config.ClientVersion
	}

	if input.Config.ClientPlatform != "" {
		m.ClientPlatform = input.Config.ClientPlatform
	}

	if input.Config.ClientArchitecture != "" {
		m.ClientArchitecture = input.Config.ClientArchitecture
	}

	if input.Config.KubectlClientVersion != "" {
		m.KubectlClientVersion = input.Config.KubectlClientVersion
	}

	var (
		vclusterCmd bytes.Buffer
		kubectlCmd  bytes.Buffer
		helmCmd     bytes.Buffer
	)

	err = vclusterTemplate.Execute(&vclusterCmd, m)
	if err != nil {
		return err
	}

	err = kubectlTemplate.Execute(&kubectlCmd, m)
	if err != nil {
		return err
	}

	err = helmTemplate.Execute(&helmCmd, m)
	if err != nil {
		return err
	}

	// Example of pulling and defining a client version for your mixin
	// fmt.Fprintf(m.Out, "\nRUN curl https://get.helm.sh/helm-%s-linux-amd64.tar.gz --output helm3.tar.gz", m.ClientVersion)
	fmt.Fprintf(m.Out, vclusterCmd.String())
	fmt.Fprintf(m.Out, kubectlCmd.String())
	fmt.Fprintf(m.Out, helmCmd.String())

	return nil
}

// validate validates that the supplied clientVersion meets the supplied semver constraint
func validate(clientVersion, constraint string) (bool, error) {
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false, errors.Wrapf(err, "unable to parse version constraint %q", constraint)
	}

	v, err := semver.NewVersion(clientVersion)
	if err != nil {
		return false, errors.Wrapf(err, "supplied client version %q cannot be parsed as semver", clientVersion)
	}

	return c.Check(v), nil
}
