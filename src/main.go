package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/css/", fs)
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
