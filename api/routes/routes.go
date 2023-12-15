package routes

import (
	"RedditAutoPost/internal/handler"

	"github.com/gin-gonic/gin"
)

// Routes configura las rutas de la aplicación en el router de Gin proporcionado.
// Define las rutas para las diferentes funciones y controladores.
//
// Parámetros:
//   - route: El router de Gin al que se agregarán las rutas.
func Routes(route *gin.Engine) {

	// Crear un grupo de rutas para versionar la API.
	redditGroup := route.Group("/api/v1")
	{
		// Definir la ruta para obtener un token de acceso.
		redditGroup.POST("/access-token", handler.GetAccessToken)

		// Actualiza el token de acceso
		redditGroup.POST("/refresh-token", handler.RefreshAccessToken)

		// Definir la ruta para crear un nuevo post.
		redditGroup.POST("/post", handler.CreatePost)
	}
}
