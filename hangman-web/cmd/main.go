package hangman_web

import (
	"fmt"
	"net/http"
)

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("./hangman-web/templates/static/stylesheets/", http.StripPrefix("./hangman-web/templates/static/stylesheets/", fs))g
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(port, server)
	if err != nil {
		return
	}
}
