package server

import (
	"RedditAutoPost/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *ginServer) error {
	router := app.server
	config := app.cfg

	apiGroup := router.Group("/api")
	{
		auth := apiGroup.Group("/auth")
		{
			auth.GET("/token", func(c *gin.Context) {
				handler.GetAccessToken(c, config)
			})

			auth.POST("/token", func(c *gin.Context) {
				handler.ObtainAccessToken(c, config)
			})

			auth.DELETE("/token", func(c *gin.Context) {
				handler.DeleteToken(c, config)
			})
		}

		posts := apiGroup.Group("/posts")
		{
			posts.POST("/create", func(c *gin.Context) {
				handler.CreatePost(c, config)
			})
		}

		credentials := apiGroup.Group("/credentials")
		{
			credentials.GET("/", handler.GetCredentials)
			credentials.POST("/create", handler.CreateCredentials)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ruta no encontrada"})
	})

	return nil
}
