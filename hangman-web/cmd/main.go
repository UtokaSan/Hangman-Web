package cmd

import (
	"fmt"
	"net/http"
)

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("./static/stylesheets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("(http://localhost:8080)")
	err := http.ListenAndServe(port, server)
	if err != nil {
		return
	}
}
