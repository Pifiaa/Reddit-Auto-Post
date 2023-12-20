package routes

import (
	"RedditAutoPost/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routes configura las rutas de la aplicación en el router de Gin proporcionado.
// Define las rutas para las diferentes funciones y controladores.
//
// Parámetros:
//   - route: El router de Gin al que se agregarán las rutas.
func Routes(route *gin.Engine) error {

	// Crear un grupo de rutas para versionar la API.
	redditGroup := route.Group("/api/v1")
	{
		// Definir la ruta para obtener un token de acceso.
		redditGroup.POST("/access-token", handler.GetAccessToken)

		// Definir la ruta para crear un nuevo post.
		redditGroup.POST("/post", handler.CreatePost)

		// Agrega credenciales de autenticación
		redditGroup.POST("/create-credentials", handler.CreateCredentials)
	}

	// Manejar el caso de ruta no encontrada.
	route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ruta no encontrada"})
	})

	return nil
}
