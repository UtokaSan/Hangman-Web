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

func Home(w http.ResponseWriter, r *http.Request) {
	test := hangman_web.Dictionnary("./words/words.txt")
	display := hangman_web.Display(test, "u")
	cookie(w, HangmanWeb{}.Word)
	fmt.Print(r.FormValue("inputText"))
	renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
		Word: display,
	})
}

// Take Value Word & save value
func cookie(w http.ResponseWriter, Word string) {
	cookie := http.Cookie{
		Name:     "Cookie",
		Value:    Word,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
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
// https://golangcode.com/add-a-http-cookie/
//https://youtu.be/ONAnstqcEcA
// *** https://www.alexedwards.net/blog/working-with-cookies-in-go ***
