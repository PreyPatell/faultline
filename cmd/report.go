package cmd

import (
	"faultline/internal/report"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Print a failure coverage report",
	RunE: func(cmd *cobra.Command, args []string) error {
		return report.Print()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
