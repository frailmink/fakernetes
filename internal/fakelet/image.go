package fakelet

import (
	"fmt"
	"net/http"
	"log/slog"
)

func RunImage(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	logger.Debug("test")

	defer func() {
		if err := r.Body.Close(); err != nil {
			logger.Error(fmt.Sprintf("Error when closing request body for running image: %v", err))
		}
	}()

	fmt.Fprint(w, "Hello")
}