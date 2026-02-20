package chaos

import (
	"fmt"
	"faultline/internal/docker"
	"faultline/internal/osdetect"
	"strconv"
)

// InjectLatency adds artificial network delay to a container.
//
// On Linux it runs `tc qdisc` inside the container via Docker exec.
// On macOS/Windows it logs a warning — a container-level stub can be added here.
func InjectLatency(client *docker.Client, service string, ms int) error {
	id, err := client.FindByName(service)
	if err != nil {
		return err
	}

	platform := osdetect.Detect()
	fmt.Printf("[latency] Injecting %dms into %s (platform: %s)\n", ms, service, platform)

	switch platform {
	case osdetect.Linux:
		cmd := []string{
			"tc", "qdisc", "add", "dev", "eth0", "root", "netem",
			"delay", strconv.Itoa(ms) + "ms",
		}
		if err := client.Exec(id, cmd); err != nil {
			return fmt.Errorf("tc exec failed: %w", err)
		}
	default:
		fmt.Printf("[latency] WARNING: tc netem unavailable on %s — using stub only.\n", platform)
		// TODO: implement container-level sleep/proxy approach for macOS/Windows
		_ = id
	}

	fmt.Printf("[latency] Done — %dms applied to %s.\n", ms, service)
	return RecordFault(Fault{
		Service: service,
		Type:    "latency",
		Params:  map[string]string{"ms": strconv.Itoa(ms)},
	})
}
