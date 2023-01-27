package hangman_web

import (
	hangman_web "Hangman/hangman-classic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./hangman-web/cmd/account.json")
	var accountUser Account
	account := string(data)
	takeMail = accountuser.Mail
	json.NewDecoder(r.Body).Decode(&accountUser)
	if strings.Contains(account, accountUser.Mail) && strings.Contains(account, accountUser.Password) {
		println("The account exist : ", accountUser.Mail)
		json.NewEncoder(w).Encode("IsAccount" + "ok")
	} else {
		println("the account not exist")
		fmt.Println(accountUser.Mail, accountUser.Password)
		json.NewEncoder(w).Encode("IsAccount :" + "ko")
	}
	if err != nil {
		fmt.Println(err)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	input := string(data)[14 : len(data)-2]
	hangmanweb.Word = wordRandom
	hangmanweb.Input = strings.ToLower(input)
	hangmanweb.InputUse = hangman_web.IsInputValid(wordRandom, hangmanweb.Input)
	if hangmanweb.Life <= 0 {
		resetGame()
	}
	if hangmanweb.InputUse == true {
		inputUseInGame = inputUseInGame + "" + hangmanweb.Input
	}
	if hangmanweb.InputUse == false && len(hangmanweb.Input) == 1 {
		if hangmanweb.Input >= string('a') && hangmanweb.Input <= string('z') || hangmanweb.Input >= string('A') && hangmanweb.Input <= string('Z') {
			hangmanweb.Life--
		}
	}
	fmt.Println("life :", hangmanweb.Life)
	hangmanweb.Display = hangman_web.Game(inputUseInGame, wordRandom)
	if hangmanweb.Display == wordRandom {
		resetGame()
	}
	err := json.NewEncoder(w).Encode(hangmanweb)
	if err != nil {
		return
	}
}

func PostDifficulty(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	difficulty = string(data)[14 : len(data)-2]
	hangmanweb.Input = difficulty
	chooseDifficulty(difficulty)
}
