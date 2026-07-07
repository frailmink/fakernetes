package fakelet

import (
	"net/http"
	"log/slog"

	"github.com/frailmink/fakernetes/internal/server"
)

func Execute(logger *slog.Logger) error {
	router := http.NewServeMux()
	router.HandleFunc("/image", server.RouteFuncWrapper([]string{http.MethodPost}, logger, RunImage))

	return server.StartServer(router, ":1000")
}
