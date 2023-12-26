package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error al iniciar configuraciones: %v", err)
	}

	ginServer := server.NewGinServer(cfg)
	err = ginServer.Start()

	if err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
