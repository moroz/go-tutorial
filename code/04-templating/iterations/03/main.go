package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type indexTemplateData struct {
	Name string
}

var indexTemplate = template.Must(template.ParseFiles("templates/index.html.tmpl"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, indexTemplateData{
		Name: "World",
	})
}

func main() {
	r := chi.NewRouter()
	r.Get("/", indexHandler)
}
