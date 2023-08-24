package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Prints pod  information",
	RunE: func(cmd *cobra.Command, args []string) error {
		klog.Info("TODO")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(podCmd)
}
