package main

import (
	"github.com/bmizerany/pat"
	"github.com/CardioLeo/go-course/pkg/config"
	"github.com/CardioLeo/go-course/pkg/handlers"
	"net/http"
)

// ran the following command to make pat package accessible
// go get github.com/bmizerany/pat
// this can now be seen in the go.mod file

func routes(app *config.AppConfig) http.Handler {
	// returns handler from http.Handler
	mux := pat.New()

	// routes will be setup here
	// routes will be taken out of main.go to be put here

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
