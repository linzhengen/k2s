package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var rootCmd = &cobra.Command{
	Use:   "k2s",
	Short: "K8s tools",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
