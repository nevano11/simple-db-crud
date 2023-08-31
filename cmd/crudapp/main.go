package main

import (
	"log"
	"simple-db-crud/internal/pkg/config"
	"simple-db-crud/internal/pkg/server"
)

func main() {
	// * Read config
	cfg, err := config.New("configs/config.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// * Create server
	apiserver, err := server.New(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	// * Run server
	if err := apiserver.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
