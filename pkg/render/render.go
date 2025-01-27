package render

import (
	"log"
        "html/template"
        "net/http"
	"path/filepath"
	"bytes"
)

// Renders templates using ...
func RenderTemplate(w http.ResponseWriter, tmpl string) {
        // create a template cache

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the requested template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	

	// he says this next part is a "belt and suspenders kind of thing to do", arbitrary
	// allows for finer grained error checking
	// helps find out as close as possible where the error is happening! awesome
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
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
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
        	        return myCache, err
		}
		matches, err := filepath.Glob("./../../templates/*.layout.tmpl")
		if err != nil {
                	return myCache, err
        	}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./../../templates/*.layout.tmpl")
		        if err != nil {
                		return myCache, err
        		}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
