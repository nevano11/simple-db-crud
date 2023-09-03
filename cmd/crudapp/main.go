package main

import (
	"github.com/spf13/viper"
	"log"
	"simple-db-crud/internal/pkg/handler"
	"simple-db-crud/internal/pkg/repository"
	"simple-db-crud/internal/pkg/server"
	"simple-db-crud/internal/pkg/service"
)

func main() {
	// Read config
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create repository, service, handler
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	// Init routes
	routes := handler.InitRoutes()

	// Create server
	apiserver, err := server.New(
		viper.GetString("logger.log-level"),
		viper.GetString("server.port"),
		routes)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Run server
	if err := apiserver.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
