package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for _, value := range r.Form {
		fmt.Printf("%s", value)
	}
	renderTemplate(w, "home")
}

func leaderboard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "leaderboard")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
