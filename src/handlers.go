package main

import (
	"net/http"
	"text/template"
)

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
