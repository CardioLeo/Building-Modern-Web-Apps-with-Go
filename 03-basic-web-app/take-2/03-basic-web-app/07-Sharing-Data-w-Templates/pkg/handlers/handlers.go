package handlers

import (
	"github.com/CardioLeo/go-course/pkg/render"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"net/http"
)

// new type for any kind of data I might want to pass, video 37
// TemplateData holds data from handlers to templates
type TemplateData struct {
	//
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	// special type for anything
	// "and in Go, when you're not sure what the datatype is, you call it an interface"
	Data map[string]interface{}
	// the following is added just for now, so it doesn't have to be added later
	// Cross-Site-Certificate-Forgery...
	CSRFToken string
	// for success message or whatever:
	Flash string
	Warning string
	Error string
}

// using the repository pattern (!!)
// allows swapping components out with minimal changes to code base

// Repo, the repository used by the handlers
// pointer to repository struct
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
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
        render.RenderTemplate(w, "home.page.tmpl", &TemplateData{})
}

// About is the page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

        render.RenderTemplate(w, "about.page.tmpl", &TemplateData{
		StringMap: stringMap,
	})
}

// now both handlers for Home and About have receivers, that is:
// the pointer to Repository just before the name
