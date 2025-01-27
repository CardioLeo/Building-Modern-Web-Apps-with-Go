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

