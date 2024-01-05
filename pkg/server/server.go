package server

import (
	"RedditAutoPost/pkg/server/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	fmt.Println("Iniciando servidor...")
	app := fiber.New()

	// Verifica si la instancia de Fiber es nula
	if app == nil {
		panic("Error al crear la instancia de Fiber")
	}

	// Configura tus rutas aquí
	routes.SetupRoutes(app)

	// Intenta iniciar el servidor en el puerto 3000
	err := app.Listen(":3000")

	// Verifica si hay algún error al intentar iniciar el servidor
	if err != nil {
		panic(fmt.Sprintf("Error al iniciar el servidor: %v", err))
	}

	fmt.Println("Servidor corriendo en el puerto 3000")
}
