package router

import (
	"github.com/aleroxac/goexpert-ratelimiter/handler"
	"github.com/go-chi/chi"
)

func InitializeRoutes(router *chi.Mux) {
	router.Get("/api/v1/healthz", handler.HealthzHandler)
	router.Get("/api/v1/zipcode/{zipcode}", handler.ZipCodeHandler)
}
