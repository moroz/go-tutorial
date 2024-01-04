package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type indexTemplateData struct {
	Name string
}

var indexTemplate = template.Must(
	template.ParseFiles("templates/index.html.tmpl"))
var contactTemplate = template.Must(
	template.ParseFiles("templates/contact.html.tmpl"))
var notFoundTemplate = template.Must(
	template.ParseFiles("templates/404.html.tmpl"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, indexTemplateData{
		Name: "World",
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contactTemplate.Execute(w, nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	notFoundTemplate.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", indexHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(notFoundHandler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
