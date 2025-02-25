package main

import (
	"fmt"
	"net/http"
)

func main() {
	// instead of just printing hello world to the screen
	// fmt.Println("Hello, world")

	http.HandleFunc("/", func(w http.ResponseWriter(), r *http.Request){
		fmt.Fprintf(w, "Hello, world!") // Fprintf requires a ResponseWriter to write to
	})
	// uniform resource locator, url, -> "/"
}
