package hangman_web

import (
	hangman_web "Hangman/hangman-classic"
	"fmt"
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
	test := hangman_web.Dictionnary("./hangman-web/words/words.txt")
	renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
		Word: test,
	})

}

func renderTemplate(w http.ResponseWriter, tmpl string, p HangmanWeb) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}

//function envoyerFormulaire(event) {
//    event.preventDefault(); // Empêche le rechargement de la page
//
//    var form = event.target; // Récupère le formulaire
//
//    var requestOptions = {
//        method: form.method,
//        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
//        body: new FormData(form)
//    };
//
//    fetch(form.action, requestOptions)
//        .then(response => response.json())
//        .then(data => console.log(data))
//        .catch(error => console.error(error))
//}

//func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
//	// read cookie
//	var cookie,err = req.Cookie("SessionID")
//	if err == nil {
//		var cookievalue = cookie.Value
//		w.Write([]byte(cookievalue))
//	}
//}
