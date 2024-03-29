package hangman_web

import (
	hangman_web "Hangman/hangman-classic"
	"fmt"
	"net/http"
)

func resetGame() {
	inputUseInGame = ""
	hangmanweb.Life = 7
}

func chooseDifficulty(input string) {
	if input == "Easy" {
		wordRandom = hangman_web.Dictionnary("./hangman-web/words/easy.txt")
	} else if input == "Moyen" {
		wordRandom = hangman_web.Dictionnary("./hangman-web/words/medium.txt")
	} else {
		wordRandom = hangman_web.Dictionnary("./hangman-web/words/hard.txt")
	}
}

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Login)
	server.HandleFunc("/difficulty", Difficulty)
	server.HandleFunc("/hangman", Home)
	server.HandleFunc("/win", Win)
	server.HandleFunc("/lose", Lose)
	server.HandleFunc("/post", Post)
	server.HandleFunc("/post-difficulty", PostDifficulty)
	server.HandleFunc("/post-login", LoginPost)
	fs := http.FileServer(http.Dir("hangman-web/templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("(http://localhost:8080) on port ", port)
	hangmanweb.Life = 7
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
