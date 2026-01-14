package main

import (
	"urlshortener.com/src/handlers"
	"urlshortener.com/src/repository"
	"urlshortener.com/src/service"
)

func main() {
	r := repository.NewInMemoryRepository()
	s := service.NewService(r)

	handlers.RegisterRoutes(s)
}
