package main

import (
	"go-ecommerce-app/config"
	_ "go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"log"
)

func main() {

	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Config file not loaded properly %v\n", err)

	}

	api.StartServer(cfg)

}
