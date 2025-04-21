package config

import (
	"html/template"
	// "log" // only include if uncommenting InfoLog below
	"github.com/alexedwards/scs/v2"
)

// App Config holds the application config
// will consist of a single struct
// later, I can put anything that I need site-wide in this file
type AppConfig struct {
	UseCache bool // this is added for Repository Pattern, video 35
	TemplateCache map[string]*template.Template
	// InfoLog added in video 34, but commented out by me because unused for now
	// InfoLog *log.Logger
	InProduction bool // bool for whether or not we are running in production or on local host;
	// need sessions to be secure in middleware.go and in scs in main.go,
	// so this abstracts away having to remember more than one place where we'll need this
	Session *scs.SessionManager
}
