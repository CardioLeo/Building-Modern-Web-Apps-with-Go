package render

import (
	"bytes"
	"log"
        "html/template"
        "net/http"
	"path/filepath"
	"github.com/CardioLeo/go-course/pkg/config"
)

// he has this but it's empty:
/*
var functions = template.FuncMap{

}
*/

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}


// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	
	var tc map[string]*template.Template
	
	// if config UseCache is true, then use what is cached!
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// if UseCache is false, then we need to make a cache
		// and so we call CreateTemplateCache
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl] // t for template
	if !ok {
		log.Fatal("could not get template from template cache")
		// this is changed from log.Fatal(err) bc err no longer exists
		// if no template, program dies
	}
	
	buf := new(bytes.Buffer) // for finer grained error checking!
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template


	/*
	// he deletes this whole section after 17 minutes, video 31
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
        err := parsedTemplate.Execute(w, nil)
        if err != nil {
                fmt.Println("error parsing template:", err) // error messages aren't supposed to begin with capital letters
                return
        }
	*/

	_, err = buf. WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
