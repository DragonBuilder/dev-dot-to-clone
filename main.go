package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")

	// http.Handle("/", r)

	slog.Info("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("Error starting server: %v", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
