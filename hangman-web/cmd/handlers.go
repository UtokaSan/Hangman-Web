package hangman_web

import (
	hangman_web "Hangman/hangman-classic"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type HangmanWeb struct {
	Word      string
	Life      int
	Display   string
	Input     string
	InputUse  bool
	LetterUse string
}

type Account struct {
	Mail     string
	Password string
}

const port = ":8080"

var test = ""
var result = hangman_web.Dictionnary("./hangman-web/words/easy.txt")
var hangmanweb HangmanWeb

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
		Word: result,
	})
}

func Difficulty(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hangman-web/templates/difficulty" + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, r)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hangman-web/templates/login" + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, r)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./hangman-web/cmd/account.json")
	var accountUser Account
	account := string(data)
	json.NewDecoder(r.Body).Decode(&accountUser)
	if strings.Contains(account, accountUser.Mail) && strings.Contains(account, accountUser.Password) {
		println("The account exist")
		json.NewEncoder(w).Encode("IsAccount :" + "ok")
	} else {
		println("the account not exist")
		json.NewEncoder(w).Encode("IsAccount :" + "ko")
	}
	if err != nil {
		fmt.Println(err)
	}
}

func Win(w http.ResponseWriter, r *http.Request) {
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
func Lose(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hangman-web/templates/lose" + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, r)
	}
}
func Post(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println("result : ", result)
	input := string(data)[14 : len(data)-2]
	fmt.Println("La valeur est :", input)
	hangmanweb.Word = result
	hangmanweb.Input = "test"
	hangmanweb.Input = input
	hangmanweb.InputUse = hangman_web.IsInputValid(result, input)
	println(hangmanweb.InputUse)
	if hangmanweb.InputUse == true {
		test = test + "" + input
		fmt.Println(test)
	} else {
		hangmanweb.Life--
	}
	fmt.Println("vie :", hangmanweb.Life)
	hangmanweb.Display = hangman_web.Game(test, result)
	fmt.Println("Display : ", hangmanweb.Display)
	err := json.NewEncoder(w).Encode(hangmanweb)
	if err != nil {
		return
	}
}

func PostDifficulty(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	input := string(data)[14 : len(data)-2]
	hangmanweb.Input = input
	fmt.Println("Difficulty : ", hangmanweb.Input)
	if input == "Easy" {
		result = hangman_web.Dictionnary("./hangman-web/words/easy.txt")
	} else if input == "Moyen" {
		result = hangman_web.Dictionnary("./hangman-web/words/medium.txt")
	} else {
		result = hangman_web.Dictionnary("./hangman-web/words/hard.txt")
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
