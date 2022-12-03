package hangman_web

import (
	"fmt"
	hangman_web "hangman-web/hangman-classic"
	"html/template"
	"net/http"
)

type HangmanWeb struct {
	Word    string
	Life    int
	Display string
}

const port = ":8080"

func Game(p *HangmanWeb) {
	c := hangman_web.HangmanData{}
	p.Word = hangman_web.Dictionnary("./words/words.txt")
	p.Display = hangman_web.Display(p.Word, c.RandomLetter)
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := HangmanWeb{}
	Game(&p)
	renderTemplate(w, "./hangman-web/templates/home", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p HangmanWeb) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
// https://www.youtube.com/watch?v=GnLHI_nekm8
