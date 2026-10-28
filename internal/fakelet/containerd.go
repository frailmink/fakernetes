package fakelet

import (
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
)

func startUpContainerd(socket string) (*containerd.Client, error) {
	// Tries connecting to the containerd socket specified in the command line
	client, err := containerd.New(socket)
	if err != nil {
		return nil, fmt.Errorf("error: failed connecting to containerd: %v", err)
	}

	return client, nil
}
