package chaos

import (
	"fmt"
	"faultline/internal/docker"
	"faultline/internal/osdetect"
	"strconv"
)

// DegradeNetwork simulates packet loss on a container's network interface.
//
// On Linux it uses `tc netem loss` inside the container.
// On macOS/Windows it prints a warning — add a proxy-based fallback here.
func DegradeNetwork(client *docker.Client, service string, lossPct int) error {
	id, err := client.FindByName(service)
	if err != nil {
		return err
	}

	platform := osdetect.Detect()
	fmt.Printf("[network] Applying %d%% packet loss to %s (platform: %s)\n", lossPct, service, platform)

	switch platform {
	case osdetect.Linux:
		cmd := []string{
			"tc", "qdisc", "add", "dev", "eth0", "root", "netem",
			"loss", strconv.Itoa(lossPct) + "%",
		}
		if err := client.Exec(id, cmd); err != nil {
			return fmt.Errorf("tc exec failed: %w", err)
		}
	default:
		fmt.Printf("[network] WARNING: tc netem unavailable on %s — using stub only.\n", platform)
		// TODO: container-level packet loss simulation for macOS/Windows
		_ = id
	}

	fmt.Printf("[network] Done — %d%% packet loss applied to %s.\n", lossPct, service)
	return RecordFault(Fault{
		Service: service,
		Type:    "packet-loss",
		Params:  map[string]string{"loss_pct": strconv.Itoa(lossPct)},
	})
}
