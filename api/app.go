package api

import (
	"RedditAutoPost/api/routes"
	"RedditAutoPost/config"
	"fmt"
)

func StartApp() {
	route := routes.SetupRouter()

	port := fmt.Sprintf(":%s", config.GetEnv("server.port"))

	fmt.Println(port)

	e := route.Run(port)

	if e != nil {
		panic(e)
	}
}
