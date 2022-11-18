package hangman_web

import (
	"html/template"
	"net/http"
)

type HangmanWeb struct {
	Word string
}

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	p := HangmanWeb{}
	p.Word = "Test"
	t, _ := template.ParseFiles("templates/home.tmpl")
	t.Execute(w, p)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
// https://www.youtube.com/watch?v=GnLHI_nekm8
