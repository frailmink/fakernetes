package server

import (
	"fmt"
	"slices"
	"net/http"
	"log/slog"
)

func StartServer(router http.Handler, url string, logger *slog.Logger) error {
	logger.Debug("Server is starting")

	return http.ListenAndServe(url, router)
}

func RouteFuncWrapper(methods []string, logger *slog.Logger, f func(http.ResponseWriter, *http.Request, *slog.Logger)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				logger.Error(fmt.Sprintf("error when closing request body for running image: %v", err))
			}
		}()

		if slices.Contains(methods, r.Method) {
			f(w, r, logger)
		} else {
			fmt.Fprintf(w, "error: method not allowed")
		}
	}
}
