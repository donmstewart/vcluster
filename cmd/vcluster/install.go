package main

import (
	"github.com/spf13/cobra"

	"github.com/donmstewart/vcluster/pkg/vcluster"
)

var (
	commandFile string
)

func buildInstallCommand(m *vcluster.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs/Creates a new Virtual Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Do something here if needed
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
