package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/server"
	"log"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	server.StartServer()
}
