package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer(workflowName, stepName, image string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	_, err = cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: image,
		Tty:   false}, nil, nil, nil, "bifrost-"+workflowName+"-"+stepName)
	if err != nil {
		return "", err
	}
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		// Ignore the error, we can't do anything
		cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{Force: true})
		return "", err
	}
	return resp.ID, nil
}
