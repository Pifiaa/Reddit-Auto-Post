package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database"
	"fmt"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	fmt.Println(config.Db)
	connection, err := database.DatabaseConnect(config)
	if err != nil {
		fmt.Printf("Error al conectar la base de datos: %v\n", err)
		return
	}
	defer connection.Close()

	// server.StartServer()
}
