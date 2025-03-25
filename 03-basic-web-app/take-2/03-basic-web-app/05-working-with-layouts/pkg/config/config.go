package config

// App Config holds the application config
// will consist of a single struct
// later, I can put anything that I need site-wide in this file
type AppConfig struct {
	// all that goes in here for now is Template Cache
	TemplateCache map[string]*template.Template
}
