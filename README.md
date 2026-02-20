# faultline

A local chaos engineering CLI for developers and QA. Injects failures into Docker containers to simulate real-world conditions.

## Install

```bash
git clone https://github.com/yourname/faultline
cd faultline
go build -o faultline .
```

Requires Docker to be running.

## Commands

```bash
faultline list services                          # list running containers
faultline kill <service>                         # kill a container
faultline inject latency <service> --ms 200      # inject 200ms latency
faultline degrade network <service> --loss 15    # 15% packet loss
faultline scenario scenarios/example.yaml        # run a YAML scenario
faultline stop                                   # stop all active faults
faultline report                                 # print fault coverage report
faultline doctor                                 # interactive TUI
```

## Scenario files

Define multi-step chaos runs in YAML:

```yaml
name: basic-resilience
description: Latency → packet loss → kill

steps:
  - action: inject_latency
    service: my-api
    params:
      ms: "200"
    wait: 5

  - action: degrade_network
    service: my-api
    params:
      loss: "15"
    wait: 3

  - action: kill
    service: my-api
```

Run with `faultline scenario <file>`.

## Structure

```
cmd/           CLI commands (one file per command)
internal/
  chaos/       fault handlers + state tracking
  docker/      Docker SDK wrapper
  osdetect/    OS detection (Linux/macOS/Windows)
  scenario/    YAML parser + step runner
  report/      fault coverage report
scenarios/     example YAML files
```

## Platform notes

- **Linux** — uses `tc netem` inside containers for real latency/packet loss
- **macOS / Windows** — falls back to container-level stubs (tc not available on host)

## Dependencies

- [cobra](https://github.com/spf13/cobra) — CLI framework
- [bubbletea](https://github.com/charmbracelet/bubbletea) — TUI (doctor command)
- [docker/docker](https://github.com/docker/docker) — Docker SDK
- [yaml.v3](https://gopkg.in/yaml.v3) — scenario parsing
