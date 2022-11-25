package cmd

import (
	"fmt"
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
	t, err := template.ParseFiles("../templates/home.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}

}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
// https://www.youtube.com/watch?v=GnLHI_nekm8
