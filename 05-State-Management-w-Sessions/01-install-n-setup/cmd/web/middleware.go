package main

import (
	"net/http"
	"fmt"
	"github.com/justinas/nosurf"
)

// he deleted WriteToConsole at the end of 05/01, but I'm keeping it
// WriteToConsole only does what it says when the page is "hit"
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println ("Hit the page")
		next.ServeHTTP(w, r)
	})
}
// every time you load the page, it runs this page and prints this out
// he says, "and that's how you write middleware"

// NoSurf, creates No Surf protection from CSRF
// adds said protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads an saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
