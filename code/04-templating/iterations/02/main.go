package main

import (
	"html/template"
	"os"
)

type indexTemplateData struct {
	Name string
}

func main() {
	tmp := template.Must(template.ParseFiles("templates/index.html.tmpl"))
	tmp.Execute(os.Stdout, indexTemplateData{
		Name: "World",
	})
}
