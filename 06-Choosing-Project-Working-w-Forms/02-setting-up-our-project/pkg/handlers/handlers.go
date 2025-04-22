package handlers

import (
	"github.com/CardioLeo/bookings/pkg/render"
	"github.com/CardioLeo/bookings/pkg/config"
	"github.com/CardioLeo/bookings/pkg/models"
	"net/http"
)

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
	remoteIP := r.RemoteAddr // native to http package
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // every time someone hits the home page, it stores their IP address within the session data
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// next two lines correlate w/ remoteIP in home handler;
	// furthermore, they'll be used in the *.page.tmpl templates
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// could do:
	// m.App.Session. --> access a whole bunch of things, basically

        render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// now both handlers for Home and About have receivers, that is:
// the pointer to Repository just before the name
