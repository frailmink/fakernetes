package server

import (
	"fmt"
	"slices"
	"net/http"
	"log/slog"
)

func StartServer(router http.Handler, url string) error {
	return http.ListenAndServe(url, router)
}

func RouteFuncWrapper(methods []string, logger *slog.Logger, f func(http.ResponseWriter, *http.Request, *slog.Logger)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(methods, r.Method) {
			f(w, r, logger)
		} else {
			fmt.Fprintf(w, "Error: method not allowed")
		}
	}
}
