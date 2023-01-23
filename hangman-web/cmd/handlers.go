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
var difficulty string

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		str := string(result[len(result)/2-1])
		renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
			Word: hangman_web.Display(result, str),
		})
	}
}

func Difficulty(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/difficulty" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
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

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./hangman-web/cmd/account.json")
	var accountUser Account
	account := string(data)
	json.NewDecoder(r.Body).Decode(&accountUser)
	if strings.Contains(account, accountUser.Mail) && strings.Contains(account, accountUser.Password) {
		println("The account exist : ", accountUser.Mail)
		json.NewEncoder(w).Encode("IsAccount" + "ok")
	} else {
		println("the account not exist")
		json.NewEncoder(w).Encode("IsAccount :" + "ko")
	}
	if err != nil {
		fmt.Println(err)
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
func Post(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println("result : ", result)
	input := string(data)[14 : len(data)-2]
	fmt.Println("La valeur est :", input)
	hangmanweb.Word = result
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
	if hangmanweb.Display == result {
		resetGame(input)
	}
	fmt.Println("Display : ", hangmanweb.Display)
	err := json.NewEncoder(w).Encode(hangmanweb)
	if err != nil {
		return
	}
}

func resetGame(input string) {
	test = ""
	test = test + "" + input
	fmt.Println(difficulty)
	chooseDifficulty(difficulty)
	hangmanweb.Life = 10
}

func PostDifficulty(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	difficulty = string(data)[14 : len(data)-2]
	hangmanweb.Input = difficulty
	chooseDifficulty(difficulty)
	fmt.Println(difficulty)
}
func chooseDifficulty(input string) {
	fmt.Println("Difficulty : ", hangmanweb.Input)
	if input == "Easy" {
		result = hangman_web.Dictionnary("./hangman-web/words/easy.txt")
	} else if input == "Moyen" {
		result = hangman_web.Dictionnary("./hangman-web/words/medium.txt")
	} else {
		result = hangman_web.Dictionnary("./hangman-web/words/hard.txt")
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
