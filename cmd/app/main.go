package main

import (
	"RedditAutoPost/config"
	"RedditAutoPost/pkg/server"
)

func main() {
	config.InitConfig()

	server.StartServer()
}
