package fakelet

import (
	"context"
	"fmt"
	"log/slog"

	containerd "github.com/containerd/containerd/v2/client"
	// "github.com/containerd/containerd/v2/pkg/namespaces"
)

type containerdAPI struct {
	client *containerd.Client
}

func startUpContainerd(socket string, logger *slog.Logger) (*containerdAPI, error) {
	logger.Debug(socket)

	// Tries connecting to the default containerd socket
	client, err := containerd.New(socket)
	if err != nil {
		return nil, fmt.Errorf("error: failed connecting to containerd: %v", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			logger.Error(fmt.Sprintf("error: failed closing containerd client: %v", err))
		}
	}()

	image, err := client.Pull(context.Background(), "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	if err != nil {
		return nil, err
	}
	logger.Debug(fmt.Sprintf("%v", image))

	return &containerdAPI{
		client: client,
	}, nil
}
