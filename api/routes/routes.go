package routes

import (
	"RedditAutoPost/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	redditGroup := route.Group("/api/v1")
	{
		redditGroup.POST("/access-token", handler.GetAccessToken)
		redditGroup.POST("/post", handler.CreatePost)
	}
}
