package main

import (
	"RedditAutoPost/api/routes"
	"RedditAutoPost/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router, err := setupApp()
	if err != nil {
		log.Fatal("Error al configurar la aplicación: ", err)
	}

	port := fmt.Sprintf(":%s", config.GetEnv("server.port"))
	addr := fmt.Sprintf("http://localhost%s", port)
	log.Printf("Iniciando el servidor en %s", addr)

	serverStarted := make(chan struct{})

	go func() {
		err = router.Run(port)
		if err != nil {
			log.Fatal("Error al iniciar el servidor: ", err)
		}
		close(serverStarted)
	}()

	<-serverStarted
	log.Printf("El servidor ha sido iniciado en %s", addr)
}

func setupApp() (*gin.Engine, error) {
	err := config.InitializeConfig()
	if err != nil {
		return nil, fmt.Errorf("Error al inicializar configuraciones: %w", err)
	}

	router := gin.Default()
	routes.Routes(router)

	return router, nil
}
