package server

import (
	"fmt"
	"slices"
	"net/http"
)

func StartServer(router http.Handler, url string) error {
	return http.ListenAndServe(url, router)
}

func RouteFuncWrapper(methods []string, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(methods, r.Method) {
			f(w, r)
		} else {
			fmt.Fprintf(w, "Error: method not allowed")
		}
	}
}
