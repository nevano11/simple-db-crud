package main

import (
	"log"
	"simple-db-crud/internal/pkg/config"
	"simple-db-crud/internal/pkg/handler"
	"simple-db-crud/internal/pkg/repository"
	"simple-db-crud/internal/pkg/server"
	"simple-db-crud/internal/pkg/service"
)

func main() {
	// Read config
	cfg, err := config.New("configs/config.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Create repository, service, handler
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	// Init routes
	routes := handler.InitRoutes()

	// Create server
	apiserver, err := server.New(cfg, routes)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Run server
	if err := apiserver.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
