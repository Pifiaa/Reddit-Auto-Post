package server

import (
	"RedditAutoPost/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *ginServer) error {
	router := app.app

	// Pasa la instancia de ginServer como valor al contexto de Gin
	router.Use(func(c *gin.Context) {
		c.Set("server", app)
		c.Next()
	})

	apiGroup := router.Group("/api")
	{
		auth := apiGroup.Group("/auth")
		{
			auth.POST("/token", handler.GetAccessToken)
		}

		posts := apiGroup.Group("/posts")
		{
			posts.POST("/", handler.CreatePost)
		}

		credentials := apiGroup.Group("/credentials")
		{
			credentials.POST("/create", handler.CreateCredentials)
		}
	}

	app.app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ruta no encontrada"})
	})

	return nil
}
