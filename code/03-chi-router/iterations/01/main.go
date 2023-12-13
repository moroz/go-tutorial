package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
