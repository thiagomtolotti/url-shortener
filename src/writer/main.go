package writer

import (
	"encoding/json"
	"net/http"
)

type JSON map[string]any

type Writer struct {
	http.ResponseWriter
}

func New(w http.ResponseWriter) *Writer {
	return &Writer{ResponseWriter: w}
}

func (w *Writer) NewJSONResponse(status int, payload JSON) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)

}

type JSONHandler func(w *Writer, r *http.Request)

func Adapt(f JSONHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := New(w)

		f(writer, r)
	}
}
