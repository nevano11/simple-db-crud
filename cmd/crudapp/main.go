package main

import (
	"fmt"
	"log"
	"simple-db-crud/internal/pkg/config"
)

func main() {
	cfg, err := config.New("config.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(cfg)
}
