package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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

