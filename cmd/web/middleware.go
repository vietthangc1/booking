package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Printed to Console")
		next.ServeHTTP(w,r)
	})
}

// Create CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: app.InProduction,
	})
	return csrfHandler
}

// Load Session and Save
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}