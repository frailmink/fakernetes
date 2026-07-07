package fakelet

import (
	"net/http"

	"github.com/frailmink/fakernetes/internal/server"
)

func Execute() error {
	router := http.NewServeMux()
	router.HandleFunc("/image", server.RouteFuncWrapper([]string{http.MethodPost}, CreateImage))

	return server.StartServer(router, ":1000")
}
