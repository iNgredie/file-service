package router

import (
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/live", probe)
	r.HandleFunc("/ready", probe)

	return r
}

func probe(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusOK)
}
