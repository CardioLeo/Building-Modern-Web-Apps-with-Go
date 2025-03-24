package render

import (
        "fmt"
        "html/template"
        "net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache

	// get requested template from cache

	// render the template


	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
        err := parsedTemplate.Execute(w, nil)
        if err != nil {
                fmt.Println("error parsing template:", err) // error messages aren't supposed to begin with capital letters
                return
        }
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	// one way you could do it
	myCache := map[string]*template.Template{}
	// make it an empty map, same functionality
	// but more common to see it this way
}
