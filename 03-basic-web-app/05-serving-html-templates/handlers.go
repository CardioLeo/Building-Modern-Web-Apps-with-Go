package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
