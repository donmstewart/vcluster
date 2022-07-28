package main

import (
	"github.com/spf13/cobra"

	"github.com/donmstewart/vcluster/pkg/vcluster"
)

func buildResumeCommand(m *vcluster.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume",
		Short: "Resumes a paused virtual cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
