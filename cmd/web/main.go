package main

import (
	"net/http"

	"github.com/pratik25sharma/firstApi/cmd/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
