package vcluster

import (
	"io/ioutil"
	"sort"
	"testing"

	"get.porter.sh/porter/pkg/exec/builder"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	testcases := []struct {
		name            string        // Test case name
		file            string        // Path to th test input yaml
		wantDescription string        // Description that you expect to be found
		wantArguments   []string      // Arguments that you expect to be found
		wantFlags       builder.Flags // Flags that you expect to be found
		wantSuffixArgs  []string      // Suffix arguments that you expect to be found
		wantSuppress    bool
	}{
		{"create", "testdata/create-input.yaml", "Create a virtual k3s cluster",
			[]string{"create", "vcluster-1", "--namespace", "host-vcluster-1"},
			nil, nil, false},
		{"connect", "testdata/connect-input.yaml", "Connect to an existing virtual k3s cluster",
			[]string{"connect", "vcluster-1", "--namespace", "host-vcluster-1", "&"},
			nil, nil, false},
		{"delete", "testdata/delete-input.yaml", "Delete a vcluster",
			[]string{"delete", "vcluster-1", "--namespace", "host-vcluster-1"},
			nil, nil, false},
		// {"delete", "testdata/build-input.yaml", "Build image",
		// 	[]string{"build"}, builder.Flags{builder.NewFlag("f", "myfile"), builder.NewFlag("t", "practice")}, []string{"/Users/myuser/Documents"}, false},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// build test
			b, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			var action Action
			err = yaml.Unmarshal(b, &action)
			require.NoError(t, err)
			require.Len(t, action.Steps, 1)

			step := action.Steps[0]
			assert.Equal(t, tc.wantDescription, step.Description)

			args := step.GetArguments()
			assert.Equal(t, tc.wantArguments, args)

			flags := step.GetFlags()
			sort.Sort(flags)
			assert.Equal(t, tc.wantFlags, flags)

			suffixArgs := step.GetSuffixArguments()
			assert.Equal(t, tc.wantSuffixArgs, suffixArgs)

			assert.Equal(t, tc.wantSuppress, step.SuppressesOutput(), "invalid suppress-output")
		})
	}
}
