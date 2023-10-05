package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func DeleteContainer(id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	err = cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}
	return nil
}
