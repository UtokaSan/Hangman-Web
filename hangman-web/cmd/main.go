package hangman_web

import (
	"fmt"
	"net/http"
)

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Login)
	server.HandleFunc("/difficulty", Difficulty)
	server.HandleFunc("/hangman", Home)
	server.HandleFunc("/win", Win)
	server.HandleFunc("/lose", Lose)
	server.HandleFunc("/post", Post)
	server.HandleFunc("/post-difficulty", PostDifficulty)
	fs := http.FileServer(http.Dir("hangman-web/templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("(http://localhost:8080) on port ", port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
