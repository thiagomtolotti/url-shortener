package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ping)
	mux.HandleFunc("/create", createURL)
	mux.HandleFunc("/{id}", getURL)

	http.ListenAndServe(":3001", mux)
}

func ping(w http.ResponseWriter, r *http.Request) {
	createMessageResponse(w, http.StatusOK, "Service is online")
}

type CreateRequest struct {
	URL string `json:"url"`
}

func createURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		createMessageResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req CreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		createMessageResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	createMessageResponse(w, http.StatusCreated, "URL was created successfully")
}

func createMessageResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		createMessageResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := r.PathValue("id")

	createMessageResponse(w, http.StatusOK, id)
}
