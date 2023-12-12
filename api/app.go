package api

import (
	"RedditAutoPost/api/router"
	"RedditAutoPost/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine = gin.Default()

func StartApp() {
	isConfigEnable, err := config.InitializeConfig()

	if !isConfigEnable {
		fmt.Println("Error:", err)
	}

	router.Router()

	port := fmt.Sprintf(":%s", config.GetEnv("server.port"))
	e := server.Run(port)

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Building RESTful API using Gin and Gorm",
		})
	})

	if e != nil {
		panic(e)
	}
}
