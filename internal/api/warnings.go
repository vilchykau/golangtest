package api

import (
	"fmt"
	"net/http"
)

func route404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprint(w, "Wrong gateway")
}
