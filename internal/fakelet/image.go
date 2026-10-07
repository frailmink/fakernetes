package fakelet

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type runImageInputs struct {
	Repository string `json:"repo"`
}

func RunImage(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	dec := json.NewDecoder(r.Body)

	var inputs runImageInputs
	if err := dec.Decode(&inputs); err != nil {
		fmt.Fprintf(w, "error, invalid json: %v", err)
		return
	}
}
