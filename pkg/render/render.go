package render

import (
	"fmt" // need to remove later
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
		log.Println("ok is set to:", ok)
		log.Println("t is set to:", t)
		log.Println("tc[tmpl] is set to:", tc[tmpl])
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
	// pages, err := filepath.Glob("./../../templates/*.page.tmpl")
	pages, err := filepath.Glob("../../03-basic-web-app/06-sharing-data-with-templates/templates/*page.tmpl")
	// log.Println("pages is set to:", pages, "and err is set to:", err)
	// log.Println("myCache is set to:", myCache)
	if err != nil {
		log.Println("The error is in CreateTemplateCache() of render.go")
		log.Println("myCache is set to:", myCache)
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("name is set to:", name)
		ts, err := template.New(name).ParseFiles(page)
		fmt.Println("ts is set to:", ts)
		if err != nil {
			log.Println("specifically, the error is in page:", name)
			log.Println("The error is in the for loop, after parsing files, of CreateTemplateCache() of render.go")
			log.Println("myCache is set to:", myCache)
			return myCache, err
		}
		// matches, err := filepath.Glob("./../../templates/*.layout.tmpl")
		matches, err := filepath.Glob("./../../03-basic-web-app/06-sharing-data-with-templates/templates/*.layout.tmpl")
		log.Println("matches is set to:", matches)
		if err != nil {
			log.Println("The error is in the for loop, after filepath.Glob(), of CreateTemplateCache() of render.go")
                	log.Println("specifically, the error is in page:", name)
			log.Println("myCache is set to:", myCache)
			return myCache, err
        	}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./../../templates/*.layout.tmpl")
			log.Println("len(matches) is greater than zero, so: ts is set to:", ts)
		        if err != nil {
                		log.Println("The error is in the for loop, after ParseGlob(), of CreateTemplateCache() of render.go")
				log.Println("specifically, the error is in page:", name)
				log.Println("myCache is set to:", myCache)
				return myCache, err
        		}
		}
		myCache[name] = ts
	}
	log.Println("everything looks fine in CreateTemplateCache()")
	return myCache, nil
}
