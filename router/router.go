package router

import (
	"github.com/aleroxac/goexpert-ratelimiter/limiter"
	"github.com/go-chi/chi"
)

func Init() {
	router := chi.NewRouter()
	rate_limiter := limiter.InitializeRateLimiters()

	InitializeMiddlewares(router, rate_limiter)
	InitializeRoutes(router)
	InitilizeServer(router)
}
