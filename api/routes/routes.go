package routes

import (
	"RedditAutoPost/api/handler"
	"RedditAutoPost/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	isConfigEnable, err := config.InitializeConfig()

	if !isConfigEnable {
		fmt.Println("Error:", err)
	}

	route := gin.Default()

	routes(route)

	return route
}

func routes(route *gin.Engine) {

	redditGroup := route.Group("/v1")
	{
		redditGroup.POST("/access_token", handler.GetAccessToken)
		redditGroup.POST("/post", handler.CreatePost)

		//redditGroup.POST("upload", handler.PopulateSubReddits)
	}
}
