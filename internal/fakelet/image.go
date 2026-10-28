package fakelet

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/cio"
	"github.com/containerd/containerd/v2/pkg/namespaces"
	"github.com/containerd/containerd/v2/pkg/oci"
)

type fakeletRoutes struct {
	contClient *containerd.Client
}

type createImageInputs struct {
	ContainerName string `json:"name"`
	Repository    string `json:"repo"`
	Namespace     string `json:"ns"`
}

func (f *fakeletRoutes) CreateImage(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	dec := json.NewDecoder(r.Body)

	var inputs createImageInputs
	if err := dec.Decode(&inputs); err != nil {
		http.Error(
			w,
			fmt.Sprintf("error, invalid json: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	ctx := namespaces.WithNamespace(context.Background(), inputs.Namespace)
	image, err := f.contClient.Pull(ctx, inputs.Repository, containerd.WithPullUnpack)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error pulling container image: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Create the container, which is the metadata object that resources are allocated and attached to
	container, err := f.contClient.NewContainer(
		ctx,
		inputs.ContainerName,
		containerd.WithImage(image),
		containerd.WithNewSnapshot(inputs.ContainerName+"-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
	)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error creating container: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Creating the task, which is the live running process on the system
	// Will start of with status CREATED
	_, err = container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error creating task: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "container created successfully")
}
