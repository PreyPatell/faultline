package cmd

import (
	"faultline/internal/chaos"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop all active fault injections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return chaos.StopAll()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
