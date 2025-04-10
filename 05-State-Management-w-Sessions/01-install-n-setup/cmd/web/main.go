package main

import (
	"fmt"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"github.com/CardioLeo/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Main is the main application function
func main() {

	var app config.AppConfig

	// initialize session with scs
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // insists that cookies are secure, or not
	// sets to false for now, because localhost isn't secure, but in production you want it set to true

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

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
