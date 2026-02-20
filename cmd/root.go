package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "faultline",
	Short: "A local chaos engineering CLI for developers and QA",
	Long: `Faultline injects failures into Docker containers to simulate
real-world failure conditions: latency, packet loss, service kills, and more.`,
}

// Execute is the entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
