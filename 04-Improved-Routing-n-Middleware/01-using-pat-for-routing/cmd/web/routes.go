package main

import (
	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"net/http"
)

// ran the following command to make pat package accessible
// go get github.com/bmizerany/pat
// this can now be seen in the go.mod file

func routes(app *config.AppConfig) http.Handler {
	// // returns handler from http.Handler

	mux := chi.NewRouter()

	// create paths for NewRouter

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
