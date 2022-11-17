package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HangmanWeb struct {
	Word         string
	RandomLetter string
}

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	var p HangmanWeb
	p.Word = Dictionnary("words/words.txt")
	p.RandomLetter = string(p.Word[len(p.Word)/2-1])
	r.ParseForm()
	for _, value := range r.Form {
		fmt.Print(value)
	}
	t, _ := template.ParseFiles("templates/home.tmpl")
	t.Execute(w, p)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
//https://www.youtube.com/watch?v=GnLHI_nekm8
