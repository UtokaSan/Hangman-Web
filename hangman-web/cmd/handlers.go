package hangman_web

import (
	"fmt"
	hangman_web "hangman-web/hangman-classic"
	"html/template"
	"net/http"
)

type HangmanWeb struct {
	Word string
	Life int
	Test string
}

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	p := HangmanWeb{}
	p.Word = "test"
	p.Test = hangman_web.Display(p.Word, "e")
	t, err := template.ParseFiles("./hangman-web/templates/home.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}

}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
// https://www.youtube.com/watch?v=GnLHI_nekm8
