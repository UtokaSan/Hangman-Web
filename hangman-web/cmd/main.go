package hangman_web

import (
	"fmt"
	"net/http"
)

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets", fs))
	fmt.Println("(http://localhost:8080) on port ", port)
	err := http.ListenAndServe(port, server)

	if err != nil {
		return
	}
}
