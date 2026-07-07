package fakelet

import (
	"fmt"
	"net/http"
)

func CreateImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}