package config

import (
	"log"
	"html/template"
)

// AppConfig holds the application config
// anything needed site-wide can be added here
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
}
