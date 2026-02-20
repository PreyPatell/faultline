package cmd

import (
	"fmt"
	"faultline/internal/docker"
	"faultline/internal/scenario"

	"github.com/spf13/cobra"
)

var scenarioCmd = &cobra.Command{
	Use:   "scenario <file>",
	Short: "Run a YAML-defined chaos scenario",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := docker.New()
		if err != nil {
			return fmt.Errorf("could not connect to Docker: %w", err)
		}
		return scenario.Run(client, args[0])
	},
}

func init() {
	rootCmd.AddCommand(scenarioCmd)
}
