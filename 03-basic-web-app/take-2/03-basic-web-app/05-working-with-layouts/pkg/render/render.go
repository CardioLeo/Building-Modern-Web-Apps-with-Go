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

	// get all of the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// not full path, only name of template
		name := filepath.Base(page)
		// .Base() returns last element of the path
		
		ts, err := template.New(name).ParseFiles(page)
		// ts stands for "templateSet", which maybe I should use instead
		// setting ts to name, then parsing the files
		if err != nil {
			return myCache, err
		}

		// look for any layouts that exist in this directory
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		// gives slice of strings with layouts
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// if matches greater than 0, deal with
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
		// stores current template set in name of map, myCache
	}

	return myCache, nil
}
