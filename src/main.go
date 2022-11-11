package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/leaderboard", leaderboard)
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
