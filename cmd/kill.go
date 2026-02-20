package cmd

import (
	"fmt"
	"faultline/internal/chaos"
	"faultline/internal/docker"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill <service>",
	Short: "Kill a running container by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := docker.New()
		if err != nil {
			return fmt.Errorf("could not connect to Docker: %w", err)
		}
		return chaos.Kill(client, args[0])
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
