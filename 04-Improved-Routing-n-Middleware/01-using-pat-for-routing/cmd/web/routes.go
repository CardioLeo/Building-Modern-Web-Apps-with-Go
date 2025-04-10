package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	mux.Use(middleware.Recoverer) // middleware is built into chi's mux

	mux.Use(WriteToConsole) // refers to function in middleware.go of pwd
	// he deletes this call to WriteToConsole after using it for demonstration, but I'm keeping it for now

	// create paths for NewRouter
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
