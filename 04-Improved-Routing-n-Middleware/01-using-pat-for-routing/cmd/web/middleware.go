package main

import (
	"net/http"
	"fmt"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println ("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// every time you load the page, it runs this page and prints this out
// he says, "and that's how you write middleware"
