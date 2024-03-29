package hangman_web

import (
	hangman_web "Hangman/hangman-classic"
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

var inputUseInGame = ""
var wordRandom = hangman_web.Dictionnary("./hangman-web/words/easy.txt")
var hangmanweb HangmanWeb
var takeMail string
var accountuser Account
var difficulty string

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		str := string(wordRandom[len(wordRandom)/2-1])
		renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
			Word: hangman_web.Display(wordRandom, str),
		})
	}
}

func Difficulty(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/difficulty" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		fmt.Println(takeMail)
		t, err := template.ParseFiles("./hangman-web/templates/difficulty" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("./hangman-web/templates/login" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}

func Win(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/win" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("./hangman-web/templates/win" + ".html")
		if hangmanweb.Display == hangmanweb.Word {
			hangmanweb.Display = "win"
		}
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}

func Lose(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/lose" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("./hangman-web/templates/lose" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./hangman-web/templates/404" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p HangmanWeb) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}
