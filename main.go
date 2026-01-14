package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there")
	})

	s := &http.Server{
		Addr: ":8080",
	}

	s.ListenAndServe()
	fmt.Println("Server is running on http://localhost:8080")
}
