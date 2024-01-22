package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/server"
	"fmt"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln("Error al cargar variables de entorno! \n", err.Error())
	}

	connection, err := database.DatabaseConnect(config)
	if err != nil {
		fmt.Printf("Error al conectar la base de datos: %v\n", err)
		return
	}
	defer connection.Close()

	// server.StartServer(config)
	server.StartServer(config)
}
