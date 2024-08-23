package main

import (
	"log"
	cfg "validator-api/pkgs/config"
	s "validator-api/server"
	"validator-api/server/routes"
)

// @Title RESTful Ethereum Validator API
// @Version 1.0
// @description This is a API for RESTful Ethereum Validator.
// @BasePath /api/v1
func main() {
	config := cfg.NewConfig()

	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables!")
	}

	server := s.NewServer(config)

	routes.ConfigureRoutes(server)
	server.Listen()
}
