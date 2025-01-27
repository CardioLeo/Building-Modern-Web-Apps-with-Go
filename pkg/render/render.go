package render

import (
        "fmt"
        "html/template"
        "net/http"
)

// Renders templates using ...
func RenderTemplate(w http.ResponseWriter, tmpl string) {
        // first of all parse the template
        parsedTemplate, _ := template.ParseFiles("./../../03-basic-web-app/05-serving-html-templates/templates/" + tmpl)
        err := parsedTemplate.Execute(w, nil)
        if err != nil {
                fmt.Println("error parsing template:", err) // error messages aren't supposed to start with capital letters
        }
}

