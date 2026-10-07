package fakelet

import (
	"net/http"
	"log/slog"

	"github.com/frailmink/fakernetes/internal/server"
)

type FakeletConfig struct {
	ContainerdSocket string
	Logger *slog.Logger
}

func Execute(config FakeletConfig) error {
	_, err := startUpContainerd(config.ContainerdSocket, config.Logger)
	if err != nil {
		return err
	}

	router := http.NewServeMux()
	router.HandleFunc("/image", server.RouteFuncWrapper([]string{http.MethodPost}, config.Logger, RunImage))

	return server.StartServer(router, ":1000", config.Logger)
}
