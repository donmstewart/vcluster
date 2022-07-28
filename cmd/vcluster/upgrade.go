package main

import (
	"github.com/spf13/cobra"

	"github.com/donmstewart/vcluster/pkg/vcluster"
)

func buildUpgradeCommand(m *vcluster.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades an existing vcluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
