package main

import (
	"RedditAutoPost/api/routes"
	"RedditAutoPost/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// main es la función principal que inicia la aplicación.
func main() {
	// setupApp configura la aplicación y devuelve el router.
	router, err := setupApp()
	if err != nil {
		log.Println("Error al configurar la aplicación: ", err)
		os.Exit(1)
	}

	// Obtener el puerto del archivo de configuración.
	port := fmt.Sprintf(":%s", config.GetEnv("server.port"))
	addr := fmt.Sprintf("http://localhost%s", port)
	log.Printf("Iniciando el servidor en %s", addr)

	// serverStarted es un canal que se cerrará cuando el servidor haya iniciado correctamente.
	serverStarted := make(chan struct{})

	// Goroutine para iniciar el servidor de manera asíncrona.
	go func() {
		err = router.Run(port)
		if err != nil {
			log.Println("Error al iniciar el servidor: ", err)
			os.Exit(1)
		}
		close(serverStarted)
	}()

	// Esperar a que el servidor haya iniciado antes de imprimir el mensaje de inicio.
	<-serverStarted
	log.Printf("El servidor ha sido iniciado en %s", addr)
}

// setupApp inicializa y configura la aplicación.
// Devuelve el router de Gin configurado con las rutas de la aplicación.
// Si hay un error durante la inicialización, devuelve nil y el error correspondiente.
func setupApp() (*gin.Engine, error) {
	// Inicializar la configuración de la aplicación.
	err := config.InitializeConfig()
	if err != nil {
		return nil, fmt.Errorf("Error al inicializar configuraciones: %w", err)
	}

	// Crear un nuevo router Gin con la configuración por defecto.
	router := gin.Default()

	// Configurar las rutas de la aplicación.
	err = routes.Routes(router)
	if err != nil {
		if gin.IsDebugging() {
			fmt.Println("Error al configurar rutas: ", err)
		}

		return nil, err
	}

	// Devolver el router configurado y sin errores.
	return router, nil
}
