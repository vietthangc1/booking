package main

import (
	"net/http"

	"github.com/vietthangc1/booking/pkg/config"
	"github.com/vietthangc1/booking/pkg/handlers"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//Using pat
// func routes(a *config.AppConfig) http.Handler {
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
// 	mux.Get("/divide", http.HandlerFunc(handlers.Repo.Divide))
// 	return mux
// }

// Using chi
func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Thử dùng middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/divide", handlers.Repo.Divide)

	return mux
}