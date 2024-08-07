package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello gamers")
	})
	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is a comment")
	})
	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "This is a comment that has the ID of %s.\n", id)
	})
	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Posted a new comment")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Printf("Starting server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
