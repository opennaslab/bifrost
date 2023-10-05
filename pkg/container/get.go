package container

import (
	"context"

	"github.com/docker/docker/client"
)

func GetContainer(id string) (string, int, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", -1, err
	}
	resp, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return "", -1, err
	}

	return resp.State.Status, resp.State.ExitCode, nil
}
