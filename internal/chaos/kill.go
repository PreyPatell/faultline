package chaos

import (
	"fmt"
	"faultline/internal/docker"
)

// Kill sends SIGKILL to the named container and records the event.
func Kill(client *docker.Client, service string) error {
	id, err := client.FindByName(service)
	if err != nil {
		return err
	}

	fmt.Printf("[kill] Sending SIGKILL to %s (%s)...\n", service, id[:12])
	if err := client.Kill(id); err != nil {
		return fmt.Errorf("kill failed: %w", err)
	}
	fmt.Printf("[kill] Container %s killed.\n", service)

	return RecordFault(Fault{
		Service: service,
		Type:    "kill",
		Params:  map[string]string{},
	})
}
