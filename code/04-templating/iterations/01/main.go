package main

import (
	"html/template"
	"os"
)

type indexTemplateData struct {
	Name string
}

func main() {
	tmp, err := template.ParseFiles("templates/index.html.tmpl")
	if err != nil {
		panic(err)
	}
	tmp.Execute(os.Stdout, indexTemplateData{
		Name: "世界",
	})
}
