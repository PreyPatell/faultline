// Package chaos is the core engine. It tracks active faults in a temp-file
// so that "stop" and "report" commands can see what is running across invocations.
package chaos

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Fault represents a single recorded fault injection.
type Fault struct {
	Service   string            `json:"service"`
	Type      string            `json:"type"`
	Params    map[string]string `json:"params"`
	StartedAt time.Time         `json:"started_at"`
}

// stateFile returns the path used to persist active faults between CLI calls.
func stateFile() string {
	return filepath.Join(os.TempDir(), "faultline-state.json")
}

// RecordFault appends a new fault to the state file.
func RecordFault(f Fault) error {
	f.StartedAt = time.Now()
	faults, err := loadFaults()
	if err != nil {
		return err
	}
	faults = append(faults, f)
	return saveFaults(faults)
}

// ActiveFaults returns all faults that have been recorded but not stopped.
func ActiveFaults() ([]Fault, error) {
	return loadFaults()
}

// StopAll prints each recorded fault and clears the state file.
// Real cleanup (e.g. removing tc rules) should be added here per fault type.
func StopAll() error {
	faults, err := loadFaults()
	if err != nil {
		return err
	}
	if len(faults) == 0 {
		fmt.Println("No active faults to stop.")
		return nil
	}
	fmt.Printf("Stopping %d active fault(s)...\n", len(faults))
	for _, f := range faults {
		fmt.Printf("  - [%s] %s on %s\n", f.Type, f.Params, f.Service)
		// TODO: call the appropriate rollback for each fault type
	}
	return os.Remove(stateFile())
}

func loadFaults() ([]Fault, error) {
	data, err := os.ReadFile(stateFile())
	if os.IsNotExist(err) {
		return []Fault{}, nil
	}
	if err != nil {
		return nil, err
	}
	var faults []Fault
	if err := json.Unmarshal(data, &faults); err != nil {
		return nil, err
	}
	return faults, nil
}

func saveFaults(faults []Fault) error {
	data, err := json.MarshalIndent(faults, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile(), data, 0o644)
}
