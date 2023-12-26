package server

import (
	"RedditAutoPost/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Start() error
}

type ginServer struct {
	app *gin.Engine
	db  *gorm.DB
	cfg *config.Config
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	return &ginServer{
		app: gin.Default(),
		db:  db,
		cfg: cfg,
	}
}

func (s *ginServer) Start() error {
	if err := SetupRoutes(s); err != nil {
		return fmt.Errorf("error al configurar las rutas: %w", err)
	}

	serverAddr := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := http.ListenAndServe(serverAddr, s.app); err != nil {
		return fmt.Errorf("error al iniciar el servidor: %w", err)
	}

	log.Printf("Iniciando el servidor en %s", s.cfg.Server.Port)

	return nil
}
