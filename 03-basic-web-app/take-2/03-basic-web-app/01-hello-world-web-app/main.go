package main

import (
	"fmt"
	"net/http"
)

// AddValues returns the sum of two integers passed into it
func addValues(x, y int) int {
        var sum int
        sum = x + y
        return sum
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

// About is the page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
	// fmt.Fprintf(w, "This is the about page")
}

// Main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
