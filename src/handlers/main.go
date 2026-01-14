package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"urlshortener.com/src/service"
	"urlshortener.com/src/writer"
)

type ApiHandler struct {
	service service.Service
}

func RegisterRoutes(s service.Service) {
	handler := ApiHandler{
		service: s,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", writer.Adapt(handler.ping))
	mux.HandleFunc("/{id}", writer.Adapt(handler.getURL))
	mux.HandleFunc("/create", writer.Adapt(handler.createURL))

	http.ListenAndServe(":3001", mux)
}

func (ah *ApiHandler) ping(w *writer.Writer, r *http.Request) {
	w.NewJSONResponse(http.StatusOK, writer.JSON{
		"message": "Service is online",
	})
}

func (ah *ApiHandler) getURL(w *writer.Writer, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodDelete {
		w.NewJSONResponse(http.StatusMethodNotAllowed, writer.JSON{"message": "Method not allowed"})
		return
	}

	if r.Method == http.MethodDelete {
		err := ah.service.DeleteURL(r.PathValue("id"))
		if err != nil {
			switch {
			case errors.Is(err, service.ErrNotFound):
				w.NewJSONResponse(http.StatusNotFound, writer.JSON{
					"message": "URL not found",
				})
			default:
				fmt.Println(err)
				w.NewJSONResponse(http.StatusInternalServerError, writer.JSON{
					"message": "Internal Server Error",
				})
			}
			return
		}

		w.NewJSONResponse(http.StatusOK, writer.JSON{"message": "URL deleted successfully"})
		return
	}

	url, err := ah.service.GetURL(r.PathValue("id"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			w.NewJSONResponse(http.StatusNotFound, writer.JSON{
				"message": "URL not found",
			})
		default:
			fmt.Println(err)
			w.NewJSONResponse(http.StatusInternalServerError, writer.JSON{
				"message": "Internal Server Error",
			})
		}

		return
	}

	http.Redirect(w, r, url, http.StatusFound)
	return
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

func (ah *ApiHandler) createURL(w *writer.Writer, r *http.Request) {
	if r.Method != http.MethodPost {
		w.NewJSONResponse(http.StatusMethodNotAllowed, writer.JSON{
			"message": "Method not allowed",
		})
		return
	}

	var req CreateURLRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.NewJSONResponse(http.StatusBadRequest, writer.JSON{
			"message": "Invalid request",
		})
		return
	}

	id := ah.service.CreateURL(req.URL)

	w.NewJSONResponse(http.StatusCreated, writer.JSON{
		"message": "URL was created successfully",
		"id":      id,
	})
}
