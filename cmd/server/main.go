package main

import (
	"github.com/aleroxac/goexpert-ratelimiter/config"
	"github.com/aleroxac/goexpert-ratelimiter/router"
)

func main() {
	config.Init()
	router.Init()
}
