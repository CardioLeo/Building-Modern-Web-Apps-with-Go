package main

import (
	"github.com/cardioleo/go-course/pkg/config"
	"github.com/cardioleo/go-course/pkg/handlers"
	"github.com/cardioleo/go-course/pkg/render"
	"fmt"
	"log"
	"net/http"
)


const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// so now, you can go to this by putting "http://localhost:8080/about" into the
	// search bar / address bar

	// "how easy it is to start a webserver that listens in go"
	// _ = http.ListenAndServe(":8080", nil) // may be preferable

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

	// starting with the underscore like that says, "if there's an error, I don't care"
	// important comment because the underscore is so confusing if you don't know!
	// "this is the basis for any web application in go" very cool

	// if you go to "http://localhost:8080/" in a web browser while
	// you're running this command, you will see hello world. (Works.)
	// Also it does print out number of bytes twice every time that
	// you refresh. He comments on how he doesn't know why and isn't
	// worried about it right now. Mine does the same thing.
}
