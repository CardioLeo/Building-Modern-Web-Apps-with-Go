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

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	// data wanted on every page will later be put here
	return td
}

// Renders templates using ...
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// tc = app.TemplateCache
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

	td = AddDefaultData(td)

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
	pages, err := filepath.Glob("../../03-basic-web-app/06-sharing-data-with-templates/templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// matches, err := filepath.Glob("./../../templates/*.layout.tmpl")
		matches, err := filepath.Glob("./../../03-basic-web-app/06-sharing-data-with-templates/templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
        	}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./../../03-basic-web-app/06-sharing-data-with-templates/templates/*.layout.tmpl")
		        if err != nil {
				return myCache, err
        		}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
