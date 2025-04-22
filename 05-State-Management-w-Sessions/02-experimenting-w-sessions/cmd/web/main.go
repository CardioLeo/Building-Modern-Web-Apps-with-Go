package main

import (
	"fmt"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"github.com/CardioLeo/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig // made global so accessible in middleware.go, for InProduction bool
var session *scs.SessionManager // variable shadowing! - until you change initialization below

// Main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	// initialize session with scs
	session = scs.New() // no longer a new assignment so that variable shadowing is done away with
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // insists that cookies are secure, or not
	// sets to false for now, because localhost isn't secure, but in production you want it set to true

	app.Session = session // sets var in config.go to the scs just initalized

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
