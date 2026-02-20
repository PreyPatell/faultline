package cmd

import (
	"fmt"
	"faultline/internal/docker"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List faultline resources",
}

var listServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List all running Docker containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := docker.New()
		if err != nil {
			return fmt.Errorf("could not connect to Docker: %w", err)
		}
		return client.ListAndPrint()
	},
}

func init() {
	listCmd.AddCommand(listServicesCmd)
	rootCmd.AddCommand(listCmd)
}
