package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
)

// Client wraps the Docker SDK so the rest of the app never touches the SDK directly.
type Client struct {
	dc *dockerclient.Client
}

// New creates a Docker client using environment variables (DOCKER_HOST, etc.).
func New() (*Client, error) {
	dc, err := dockerclient.NewClientWithOpts(
		dockerclient.FromEnv,
		dockerclient.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	return &Client{dc: dc}, nil
}

// ListAndPrint fetches running containers and prints a summary table.
func (c *Client) ListAndPrint() error {
	containers, err := c.dc.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}
	if len(containers) == 0 {
		fmt.Println("No running containers found.")
		return nil
	}
	fmt.Printf("%-14s %-28s %-16s %s\n", "ID", "NAME", "STATUS", "IMAGE")
	fmt.Println(strings.Repeat("-", 80))
	for _, ctr := range containers {
		name := nameOf(ctr.Names)
		fmt.Printf("%-14s %-28s %-16s %s\n", ctr.ID[:12], name, ctr.Status, ctr.Image)
	}
	return nil
}

// FindByName resolves a container name to its full ID.
func (c *Client) FindByName(name string) (string, error) {
	containers, err := c.dc.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return "", err
	}
	for _, ctr := range containers {
		for _, n := range ctr.Names {
			if strings.TrimPrefix(n, "/") == name {
				return ctr.ID, nil
			}
		}
	}
	return "", fmt.Errorf("service %q not found — is the container running?", name)
}

// Kill sends SIGKILL to the container with the given ID.
func (c *Client) Kill(id string) error {
	return c.dc.ContainerKill(context.Background(), id, "SIGKILL")
}

// Exec runs a shell command inside a container (fire-and-forget).
func (c *Client) Exec(id string, cmd []string) error {
	resp, err := c.dc.ContainerExecCreate(context.Background(), id, types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return err
	}
	return c.dc.ContainerExecStart(context.Background(), resp.ID, types.ExecStartCheck{})
}

func nameOf(names []string) string {
	if len(names) == 0 {
		return "<unnamed>"
	}
	return strings.TrimPrefix(names[0], "/")
}
