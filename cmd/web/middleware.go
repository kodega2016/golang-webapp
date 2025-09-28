package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf add CSRF protection to all the POST request.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and save the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
