package vcluster

import (
	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var _ builder.ExecutableAction = Action{}
var _ builder.BuildableAction = Action{}

type Action struct {
	Name  string
	Steps []VClusterStep // using UnmarshalYAML so that we don't need a custom type per action
}

// MarshalYAML converts the action back to a YAML representation
// install:
//   vcluster:
//     ...
func (a Action) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{a.Name: a.Steps}, nil
}

// MakeSteps builds a slice of VClusterStep for data to be unmarshaled into.
func (a Action) MakeSteps() interface{} {
	return &[]VClusterStep{}
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - vcluster: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, a)
	if err != nil {
		return err
	}

	for actionName, action := range results {
		a.Name = actionName

		for _, result := range action {
			step := result.(*[]VClusterStep)
			a.Steps = append(a.Steps, *step...)
		}

		break // There is only 1 action
	}

	return nil
}

func (a Action) GetSteps() []builder.ExecutableStep {
	// Go doesn't have generics, nothing to see here...
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}

type VClusterStep struct {
	Description string
	VClusterCommand
}

// UnmarshalYAML takes any yaml in this form
// docker:
//   description: something
//   COMMAND: # pull/build/run/... -> make the PullCommand/BuildCommand/RunCommand for us
func (s *VClusterStep) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Turn the yaml into a raw map so we can iterate over the values and
	// look for which command was used
	stepMap := map[string]map[string]interface{}{}
	err := unmarshal(&stepMap)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal yaml into a raw docker command")
	}

	// get at the values defined under "vcluster"
	vclusterStep := stepMap["vcluster"]

	// Turn each command into its typed data structure
	for key, value := range vclusterStep {
		var cmd VClusterCommand

		switch key {
		case "description":
			s.Description = value.(string)

			continue
		case "create":
			cmd = &CreateCommand{}
		case "connect":
			cmd = &ConnectCommand{}
		case "disconnect":
			cmd = &DisconnectCommand{}
		case "delete":
			cmd = &DeleteCommand{}
		case "pause":
			cmd = &PauseCommand{}
		case "resume":
			cmd = &ResumeCommand{}
		case "upgrade":
			cmd = &UpgradeCommand{}
		default:
			return errors.Errorf("unsupported vcluster mixin command %s", key)
		}

		b, err := yaml.Marshal(value)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(b, cmd)
		if err != nil {
			return err
		}

		s.VClusterCommand = cmd
	}

	return nil
}

type VClusterCommand interface {
	builder.ExecutableStep
	builder.HasOrderedArguments
	builder.SuppressesOutput
}

var _ VClusterCommand = Step{}

// Actions is a set of actions, and the steps, passed from Porter.
type Actions []Action

// UnmarshalYAML takes chunks of a porter.yaml file associated with this mixin
// and populates it on the current action set.
// install:
//   vcluster:
//     ...
//   vcluster:
//     ...
// upgrade:
//   vcluster:
//     ...
func (a *Actions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, Action{})
	if err != nil {
		return err
	}

	for actionName, action := range results {
		for _, result := range action {
			s := result.(*[]VClusterStep)
			*a = append(*a, Action{
				Name:  actionName,
				Steps: *s,
			})
		}
	}
	return nil
}

type Step struct {
	Name            string        `yaml:"name"`
	Description     string        `yaml:"description"`
	KubeContext     string        `yaml:"kubeContext"`
	Arguments       []string      `yaml:"arguments,omitempty"`
	SuffixArguments []string      `yaml:"suffix-arguments,omitempty"`
	Flags           builder.Flags `yaml:"flags,omitempty"`
	Outputs         []Output      `yaml:"outputs,omitempty"`
	SuppressOutput  bool          `yaml:"suppress-output,omitempty"`
}

func (s Step) GetCommand() string {
	return "vcluster"
}

func (s Step) GetArguments() []string {
	return s.Arguments
}

func (s Step) GetSuffixArguments() []string {
	return s.SuffixArguments
}

func (s Step) GetFlags() builder.Flags {
	return s.Flags
}

func (s Step) SuppressesOutput() bool {
	return s.SuppressOutput
}

func (s Step) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}

var _ builder.OutputJsonPath = Output{}
var _ builder.OutputFile = Output{}
var _ builder.OutputRegex = Output{}

type Output struct {
	Name string `yaml:"name"`

	// See https://porter.sh/mixins/exec/#outputs
	// TODO: If your mixin doesn't support these output types, you can remove these and the interface assertions above, and from #/definitions/outputs in schema.json
	JsonPath string `yaml:"jsonPath,omitempty"`
	FilePath string `yaml:"path,omitempty"`
	Regex    string `yaml:"regex,omitempty"`
}

func (o Output) GetName() string {
	return o.Name
}

func (o Output) GetJsonPath() string {
	return o.JsonPath
}

func (o Output) GetFilePath() string {
	return o.FilePath
}

func (o Output) GetRegex() string {
	return o.Regex
}
