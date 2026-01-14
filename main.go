package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ping)

	http.ListenAndServe(":3001", mux)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Service is online",
	})
}
