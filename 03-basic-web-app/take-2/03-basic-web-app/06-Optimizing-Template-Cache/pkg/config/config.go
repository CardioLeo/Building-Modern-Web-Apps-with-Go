package config

import (
	"html/template"
)

// App Config holds the application config
// will consist of a single struct
// later, I can put anything that I need site-wide in this file
type AppConfig struct {
	// he adds this in video 34 (or 35, but made 34 retroactively?)
	// UseCache bool
	TemplateCache map[string]*template.Template
	// InfoLog added in video 34, but commented out by me because unused for now
	// InfoLog *log.Logger
}
