package server

/*import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/server/routes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/schollz/progressbar/v3"
)

func StartServer(config config.Config) {
	// Crear un canal para gestionar las señales de apagado
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Inicializar la barra de progreso
	bar := initializeProgressBar()

	// Iniciar el servidor Fiber
	app := startFiberServer(config)

	// Actualizar la barra de progreso hasta que el servidor esté listo
	updateProgressBar(bar)

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] Servidor Fiber inicializado!", currentTime)

	// Bloquear hasta que se apague el servidor
	<-shutdown

	// Apagar correctamente el servidor Fiber
	shutdownFiberServer(app)
}

func initializeProgressBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("[cyan][reset] Iniciando servidor..."),
	)
}

func startFiberServer(config config.Config) *fiber.App {
	app := fiber.New()
	micro := fiber.New()

	// Verifica si la instancia de Fiber es nula
	if app == nil {
		panic("Error al crear la instancia de Fiber")
	}

	// Middleware personalizado para el seguimiento de rutas
	trackRoutes(app)

	// Configura de rutas
	routes.SetupRoutes(app, micro)

	// Ejecutar el servidor Fiber en una goroutine
	go func() {
		port := fmt.Sprintf(":%s", config.Server.Port)

		// Intenta iniciar el servidor en el puerto 3000
		err := app.Listen(port)

		// Verifica si hay algún error al intentar iniciar el servidor
		if err != nil {
			panic(fmt.Sprintf("Error al iniciar el servidor: %v", err))
		}
	}()

	return app
}

func trackRoutes(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("%s %s\n", c.Method(), c.Path())
		return c.Next()
	})
}

func updateProgressBar(bar *progressbar.ProgressBar) {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond) // Simular el retraso en el arranque del servidor
		bar.Add(1)
	}
}

func shutdownFiberServer(app *fiber.App) {
	// Apagar correctamente el servidor Fiber
	err := app.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nServidor cerrado...")
}
*/
