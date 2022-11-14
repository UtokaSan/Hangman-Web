package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("templates/css/", fs))
	http.HandleFunc("/leaderboard", leaderboard)
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
