package main

import (
	"encoding/json"
	"net/http"

	"urlshortener.com/src/writer"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", writer.Adapt(ping))
	mux.HandleFunc("/create", writer.Adapt(createURL))
	mux.HandleFunc("/{id}", writer.Adapt(getURL))

	http.ListenAndServe(":3001", mux)
}

func ping(w *writer.Writer, r *http.Request) {
	w.NewJSONResponse(http.StatusOK, writer.JSON{
		"message": "Service is online",
	})
}

type CreateRequest struct {
	URL string `json:"url"`
}

func createURL(w *writer.Writer, r *http.Request) {
	if r.Method != http.MethodPost {
		w.NewJSONResponse(http.StatusMethodNotAllowed, writer.JSON{
			"message": "Method not allowed",
		})
		return
	}

	var req CreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.NewJSONResponse(http.StatusBadRequest, writer.JSON{"message": "Invalid request"})
		return
	}

	w.NewJSONResponse(http.StatusCreated, writer.JSON{"message": "URL was created successfully"})
}

func getURL(w *writer.Writer, r *http.Request) {
	if r.Method != http.MethodGet {
		w.NewJSONResponse(http.StatusMethodNotAllowed, writer.JSON{"message": "Method not allowed"})
		return
	}

	id := r.PathValue("id")

	w.NewJSONResponse(http.StatusOK, writer.JSON{"id": id})
}
