package handlers

import (
	"github.com/cardioleo/go-course/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
        render.RenderTemplate(w, "home.page.tmpl")
}

// first thing that should appear before function is the function name,
// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
        render.RenderTemplate(w, "about.page.tmpl")
}
