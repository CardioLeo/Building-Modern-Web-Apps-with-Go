package handlers

import (
	"github.com/cardioleo/go-course/pkg/render"
	"github.com/cardioleo/go-course/pkg/config"
	"net/http"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{} //special data
					// "and in Go, when you're not sure
					// what the data type is, we call it
					// an interface."
	CSRFtoken string
	Flash string
	Warning string
	Error string
}

// the repository pattern
// allows swapping stuff out of code base with minimal changes

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
        render.RenderTemplate(w, "home.page.tmpl", &TemplateData{})
}

// first thing that should appear before function is the function name,
// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."


	// send data to the template (no way to do this as of right now)
	// make changes in render template
	render.RenderTemplate(w, "about.page.tmpl", &TemplateData{})
}
