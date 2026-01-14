package main

import (
	"net/http"

	"urlshortener.com/src/handlers"
	"urlshortener.com/src/writer"
)

func main() {
	mux := http.NewServeMux()

	for path, handler := range handlers.Handlers {
		mux.HandleFunc(path, writer.Adapt(handler))
	}

	http.ListenAndServe(":3001", mux)
}
