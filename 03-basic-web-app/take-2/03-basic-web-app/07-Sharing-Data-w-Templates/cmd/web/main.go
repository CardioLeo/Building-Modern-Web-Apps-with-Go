package main

import (
	"fmt"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"github.com/CardioLeo/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Main is the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // added in video 35

	// creates repository variable
	repo := handlers.NewRepo(&app)
	// pass back to handlers
	// to use repository pattern
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // nil is for handler
	// the underscore is for saying that if there's an error, we don't care
	// because ListenAndServe does return an error, I guess
}
