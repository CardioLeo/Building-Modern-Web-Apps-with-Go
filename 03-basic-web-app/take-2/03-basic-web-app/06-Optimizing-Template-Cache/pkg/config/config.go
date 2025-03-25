package config

import (
	"html/template"
)

// App Config holds the application config
// will consist of a single struct
// later, I can put anything that I need site-wide in this file
type AppConfig struct {
	UseCache bool // this is added for Repository Pattern, video 35
	TemplateCache map[string]*template.Template
	// InfoLog added in video 34, but commented out by me because unused for now
	// InfoLog *log.Logger
}
