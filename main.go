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
	NewJSONResponse(w, http.StatusOK, JSON{
		"message": "Service is online",
	})
}

type CreateRequest struct {
	URL string `json:"url"`
}

func createURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		NewJSONResponse(w, http.StatusMethodNotAllowed, JSON{
			"message": "Method not allowed",
		})
		return
	}

	var req CreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		NewJSONResponse(w, http.StatusBadRequest, JSON{"message": "Invalid request"})
		return
	}

	NewJSONResponse(w, http.StatusCreated, JSON{"message": "URL was created successfully"})
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		NewJSONResponse(w, http.StatusMethodNotAllowed, JSON{"message": "Method not allowed"})
		return
	}

	id := r.PathValue("id")

	NewJSONResponse(w, http.StatusOK, JSON{"id": id})
}

type JSON map[string]any

func NewJSONResponse(w http.ResponseWriter, status int, res JSON) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
