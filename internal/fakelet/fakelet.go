package fakelet

import (
	"net/http"
	"log/slog"
	"fmt"

	"github.com/frailmink/fakernetes/internal/server"
)

type FakeletConfig struct {
	ContainerdSocket string
	Logger *slog.Logger
}

func Execute(config FakeletConfig) error {
	contClient, err := startUpContainerd(config.ContainerdSocket)
	if err != nil {
		return err
	}

	defer func() {
		if err := contClient.Close(); err != nil {
			config.Logger.Error(fmt.Sprintf("error: failed closing containerd client: %v", err))
		}
	}()

	// This object is used to pass in specific values to each route
	fr := fakeletRoutes{
		contClient: contClient,
	}

	router := http.NewServeMux()
	router.HandleFunc("/image", server.RouteFuncWrapper([]string{http.MethodPost}, config.Logger, fr.CreateImage))

	return server.StartServer(router, ":1000", config.Logger)
}
