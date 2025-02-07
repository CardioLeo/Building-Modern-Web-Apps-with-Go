package render

import (
	"log"
        "html/template"
        "net/http"
	"path/filepath"
	"bytes"
	"github.com/cardioleo/go-course/pkg/config"
	"github.com/cardioleo/go-course/pkg/models"
)

var functions = template.FuncMap{

} // okay, good news, this is still blank in his code for video 37.

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Renders templates using ...
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// line above now creates a template cache

	// get the requested template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Println("the error is in RenderTemplate in render.go")
		log.Fatal("could not get template from template cache")
	}
	

	// he says this next part is a "belt and suspenders kind of thing to do", arbitrary
	// allows for finer grained error checking
	// helps find out as close as possible where the error is happening! awesome
	buf := new(bytes.Buffer)

	// was err instead of _
	_ = t.Execute(buf, td) // td is the template data passed in from models
	// he gets rid of this next part
	/*
	if err != nil {
		log.Println(err)
	}
	*/

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)	
	myCache := map[string]*template.Template{} // exactly the same functionality as using make
	// get all of the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob("./../../templates/*.page.tmpl")
	if err != nil {
		log.Println("The error is in CreateTemplateCache() of render.go")
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("The error is in the for loop, after parsing files, of CreateTemplateCache() of render.go")
        	        return myCache, err
		}
		matches, err := filepath.Glob("./../../templates/*.layout.tmpl")
		if err != nil {
			log.Println("The error is in the for loop, after filepath.Glob(), of CreateTemplateCache() of render.go")
                	return myCache, err
        	}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./../../templates/*.layout.tmpl")
		        if err != nil {
                		log.Println("The error is in the for loop, after ParseGlob(), of CreateTemplateCache() of render.go")
				return myCache, err
        		}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
