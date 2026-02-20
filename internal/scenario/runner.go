package scenario

import (
	"fmt"
	"faultline/internal/chaos"
	"faultline/internal/docker"
	"strconv"
	"time"
)

// Run parses a scenario file and executes each step in order.
func Run(client *docker.Client, file string) error {
	sc, err := Parse(file)
	if err != nil {
		return fmt.Errorf("failed to parse scenario %q: %w", file, err)
	}

	fmt.Printf("=== Scenario: %s ===\n%s\n\n", sc.Name, sc.Description)

	for i, step := range sc.Steps {
		fmt.Printf("[%d/%d] %s → %s  params=%v\n", i+1, len(sc.Steps), step.Action, step.Service, step.Params)

		if err := dispatch(client, step); err != nil {
			// Log the error but keep going so remaining steps still run.
			fmt.Printf("[%d/%d] ERROR: %v\n", i+1, len(sc.Steps), err)
		}

		if step.Wait > 0 {
			fmt.Printf("[%d/%d] Waiting %ds...\n", i+1, len(sc.Steps), step.Wait)
			time.Sleep(time.Duration(step.Wait) * time.Second)
		}
	}

	fmt.Println("\nScenario complete.")
	return nil
}

// dispatch maps a step action string to the correct chaos handler.
func dispatch(client *docker.Client, step Step) error {
	switch step.Action {
	case "kill":
		return chaos.Kill(client, step.Service)

	case "inject_latency":
		ms, _ := strconv.Atoi(step.Params["ms"])
		return chaos.InjectLatency(client, step.Service, ms)

	case "degrade_network":
		loss, _ := strconv.Atoi(step.Params["loss"])
		return chaos.DegradeNetwork(client, step.Service, loss)

	default:
		return fmt.Errorf("unknown action %q", step.Action)
	}
}
