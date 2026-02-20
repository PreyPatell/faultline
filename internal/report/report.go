// Package report prints a summary of all recorded fault injections.
package report

import (
	"fmt"
	"faultline/internal/chaos"
	"strings"
)

// Print reads the current state file and outputs a formatted report.
func Print() error {
	faults, err := chaos.ActiveFaults()
	if err != nil {
		return err
	}

	fmt.Println("=== Faultline Report ===")
	fmt.Printf("Total faults recorded: %d\n\n", len(faults))

	if len(faults) == 0 {
		fmt.Println("No faults recorded yet. Run some chaos first!")
		return nil
	}

	fmt.Printf("%-20s %-14s %-10s %s\n", "SERVICE", "TYPE", "STARTED", "PARAMS")
	fmt.Println(strings.Repeat("-", 70))

	for _, f := range faults {
		params := formatParams(f.Params)
		fmt.Printf("%-20s %-14s %-10s %s\n",
			f.Service, f.Type, f.StartedAt.Format("15:04:05"), params)
	}

	return nil
}

func formatParams(p map[string]string) string {
	var parts []string
	for k, v := range p {
		parts = append(parts, k+"="+v)
	}
	return strings.Join(parts, " ")
}
