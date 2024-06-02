package router

import (
	"github.com/aleroxac/goexpert-ratelimiter/limiter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitializeMiddlewares(router *chi.Mux, limiter *limiter.RateLimiter) {
	router.Use(limiter.Middleware)

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
}
