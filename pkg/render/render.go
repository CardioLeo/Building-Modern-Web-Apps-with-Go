package render

import (
        "fmt"
	"log"
        "html/template"
        "net/http"
)

// Renders templates using ...
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
        // first of all parse the template
        parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
        // ./../../03-basic-web-app/05-serving-html-templates/
	err := parsedTemplate.Execute(w, nil)
        if err != nil {
                fmt.Println("error parsing template:", err) // error messages aren't supposed to start with capital letters
        }
}

// tc for template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t] // if inMap == false, then template not in cache
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in cache
		log.Println("using cached template")
	}
	
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		// path from render pkg: ./../../03-basic-web-app/05-serving-html-templates/
		fmt.Sprintf("./../../templates/%s", t),
		"./../../templates/base.layout.tmpl", // path from main to template dir
	}

	// parse the template

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl

	return nil
}
