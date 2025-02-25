package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the about page")
}

func AddValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func main() {
	http.HandlerFunc("/", Home)
	http.HandlerFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
