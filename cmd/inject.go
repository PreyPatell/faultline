package cmd

import (
	"fmt"
	"faultline/internal/chaos"
	"faultline/internal/docker"

	"github.com/spf13/cobra"
)

var latencyMs int

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject a fault into a service",
}

var injectLatencyCmd = &cobra.Command{
	Use:   "latency <service>",
	Short: "Inject network latency into a container",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := docker.New()
		if err != nil {
			return fmt.Errorf("could not connect to Docker: %w", err)
		}
		return chaos.InjectLatency(client, args[0], latencyMs)
	},
}

func init() {
	injectLatencyCmd.Flags().IntVar(&latencyMs, "ms", 100, "Latency to inject in milliseconds")
	injectCmd.AddCommand(injectLatencyCmd)
	rootCmd.AddCommand(injectCmd)
}
