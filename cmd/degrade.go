package cmd

import (
	"fmt"
	"faultline/internal/chaos"
	"faultline/internal/docker"

	"github.com/spf13/cobra"
)

var packetLoss int

var degradeCmd = &cobra.Command{
	Use:   "degrade",
	Short: "Degrade a network aspect of a service",
}

var degradeNetworkCmd = &cobra.Command{
	Use:   "network <service>",
	Short: "Simulate packet loss on a container",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := docker.New()
		if err != nil {
			return fmt.Errorf("could not connect to Docker: %w", err)
		}
		return chaos.DegradeNetwork(client, args[0], packetLoss)
	},
}

func init() {
	degradeNetworkCmd.Flags().IntVar(&packetLoss, "loss", 10, "Packet loss percentage (0-100)")
	degradeCmd.AddCommand(degradeNetworkCmd)
	rootCmd.AddCommand(degradeCmd)
}
