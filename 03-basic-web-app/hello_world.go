package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello, world")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// first parameter is the url that I want this to listen to
		n, err := fmt.Fprintf(w, "Hello, world!")
		// fmt.Fprintf(w, "Hello, world!")
	})
}
