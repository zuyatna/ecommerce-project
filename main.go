package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthCheckHandler)

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
