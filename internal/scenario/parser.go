// Package scenario handles loading and validating YAML chaos scenario files.
package scenario

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Step is one action inside a scenario.
type Step struct {
	Action  string            `yaml:"action"`  // kill | inject_latency | degrade_network
	Service string            `yaml:"service"` // container name
	Params  map[string]string `yaml:"params"`  // action-specific key/value pairs
	Wait    int               `yaml:"wait"`    // seconds to pause after this step
}

// Scenario is the top-level structure of a .yaml chaos file.
type Scenario struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
}

// Parse reads and deserialises a scenario YAML file.
func Parse(file string) (*Scenario, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var sc Scenario
	if err := yaml.Unmarshal(data, &sc); err != nil {
		return nil, err
	}
	return &sc, nil
}
