package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const portNumber = ":8080"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err) // error messages aren't supposed to begin with capital letters
		return
	}
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.page.tmpl")
}

// About is the page handler
func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about.page.tmpl")
}

// Main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
