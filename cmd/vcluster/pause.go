package main

import (
	"github.com/spf13/cobra"

	"github.com/donmstewart/vcluster/pkg/vcluster"
)

func buildPauseCommand(m *vcluster.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Pauses a running virtual cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
