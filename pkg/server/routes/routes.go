package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, micro *fiber.App) {

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	/*micro.Route("/auth", func(router fiber.Router) {
		router.Post("/", handler.CreateToken())
		// router.Get("/", handler.GetAccessToken())
	})*/

	/*micro.Route("/post", func(router fiber.Router) {
		router.Post("/")
		router.Post("/")
		router.Post("/")
	})*/

	micro.Route("/credential", func(router fiber.Router) {
		router.Post("/", handler.)
	})
}
