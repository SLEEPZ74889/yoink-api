package router

import (
	"github.com/SLEEPZ74889/yoink-api/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(linkHandler *handler.LinkHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handler.HealthHandler)
	r.Post("/links", linkHandler.Create)

	return r
}
