package hangman_web

import (
	"fmt"
	"net/http"
)

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Home)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates"))))
	fmt.Println("(http://localhost:8080) - Server started on port", port)

	err := http.ListenAndServe(port, server)
	if err != nil {
		return
	}
}
