package main

import (
	"fmt"
	"net/http"
)

func main() {
	func Home(w http.ResponseWriter, r *http.Request){
		
	}


	// uniform resource locator, url, -> "/"
	_ = http.ListenAndServe(":8080", nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
