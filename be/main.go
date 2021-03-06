package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/health", HealthCheckHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
