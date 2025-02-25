package main

import (
	"fmt"
	"net/http"
)

func main() {
	// instead of just printing hello world to the screen
	// fmt.Println("Hello, world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, world!")
		// Fprintf requires a ResponseWriter to write to
		// returns an int and an error
		// fmt.Println("Bytes written:", + n) // doesn't work for type int
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})
	// uniform resource locator, url, -> "/"
	_ = http.ListenAndServe(":8080", nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
