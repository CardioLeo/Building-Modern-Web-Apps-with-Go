package config

import (
	"html/template"
)

// AppConfig holds the application config
// anything needed site-wide can be added here
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
