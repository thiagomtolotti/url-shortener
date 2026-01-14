package handlers

import (
	"encoding/json"
	"net/http"

	"urlshortener.com/src/service"
	"urlshortener.com/src/writer"
)

var Handlers map[string]writer.JSONHandler = map[string]writer.JSONHandler{
	"/":       ping,
	"/{id}":   getURL,
	"/create": createURL,
}

func ping(w *writer.Writer, r *http.Request) {
	w.NewJSONResponse(http.StatusOK, writer.JSON{
		"message": "Service is online",
	})
}

func getURL(w *writer.Writer, r *http.Request) {
	if r.Method != http.MethodGet {
		w.NewJSONResponse(http.StatusMethodNotAllowed, writer.JSON{"message": "Method not allowed"})
		return
	}

	w.NewJSONResponse(http.StatusOK, writer.JSON{"url": ""})
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

	id := service.CreateURL(req.URL)

	w.NewJSONResponse(http.StatusCreated, writer.JSON{
		"message": "URL was created successfully", "id": id,
	})
}
