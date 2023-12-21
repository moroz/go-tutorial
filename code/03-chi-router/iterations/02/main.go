package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<h1>Welcome to Wang Xiaoming's Website!</h1>
		<a href="/contact">Contact</a>
	`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<h1>Contact me</h1>
		<a href="/">Back to homepage</a>
	`)
}
