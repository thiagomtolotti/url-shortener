package main

import (
	"urlshortener.com/src/handlers"
	"urlshortener.com/src/infra"
	"urlshortener.com/src/repository"
	"urlshortener.com/src/service"
)

func main() {
	err := infra.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	db, err := infra.ConnectToDB()
	if err != nil {
		panic(err)
	}

	r := repository.NewSQLRepository(db)
	s := service.NewService(r)

	handlers.RegisterRoutes(s)
}
