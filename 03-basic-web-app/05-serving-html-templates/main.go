package main

import (
	"fmt"
	"net/http"
	"html/template"
)


const portNumber = ":8080"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// first of all parse the template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err) // error messages aren't supposed to start with capital letters
	}
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// first thing that should appear before function is the function name,
// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}
// main is the main application function
func main() {
	// fmt.Println("Hello, world")
	/*
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
	*/
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// so now, you can go to this by putting "http://localhost:8080/about" into the
	// search bar / address bar

	// "how easy it is to start a webserver that listens in go"
	// _ = http.ListenAndServe(":8080", nil) // may be preferable

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

	// starting with the underscore like that says, "if there's an error, I don't care"
	// important comment because the underscore is so confusing if you don't know!
	// "this is the basis for any web application in go" very cool

	// if you go to "http://localhost:8080/" in a web browser while
	// you're running this command, you will see hello world. (Works.)
	// Also it does print out number of bytes twice every time that
	// you refresh. He comments on how he doesn't know why and isn't
	// worried about it right now. Mine does the same thing.
}
