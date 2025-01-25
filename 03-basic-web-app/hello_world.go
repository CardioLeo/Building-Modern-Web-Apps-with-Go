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
		// fmt.Println("Bytes written:" + n)
		// "n is an int, and you can't print an int like that" - dude is being Dr Seuss
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})


	// "how easy it is to start a webserver that listens in go"
	http.ListenAndServe(":8080", nil)
}
