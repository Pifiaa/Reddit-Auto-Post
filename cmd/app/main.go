package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error al iniciar configuraciones: %v", err)
	}

	db, err := database.DatabaseConnect(cfg)
	if err != nil {
		log.Printf("Error al conectar a la base de datos: %v", err)
		return
	}

	defer db.Close()

	ginServer := server.NewGinServer(cfg, db.GetDb())

	err = ginServer.Start()
	if err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
