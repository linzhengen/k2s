package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Prints node information",
	RunE: func(cmd *cobra.Command, args []string) error {
		klog.Info("TODO")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
