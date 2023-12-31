package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/database/models/token"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// TokenService maneja las operaciones relacionadas con tokens.
type TokenService struct {
	db database.Database
}

// NewTokenService crea una nueva instancia de TokenService.
func NewTokenService() (*TokenService, error) {
	db, err := database.DatabaseConnect()
	if err != nil {
		return nil, fmt.Errorf("Error al conectar la base de datos: %w", err)
	}
	return &TokenService{db: db}, nil
}

// CreateAccessToken crea un nuevo token de acceso.
func (ts *TokenService) CreateAccessToken(accessToken string, expiration time.Time) error {
	ts.DeleteToken()

	newToken := token.Tokens{Token: accessToken, Expiration: expiration}

	result := ts.db.GetDb().Where(token.Tokens{Token: accessToken}).Assign(&newToken).FirstOrCreate(&newToken)
	if result.Error != nil {
		return fmt.Errorf("Error al crear el token de acceso: %w", result.Error)
	}

	log.Println("Token creado exitosamente")
	return nil
}

// GetToken obtiene el token almacenado en la base de datos.
func (ts *TokenService) GetToken() (token.Tokens, error) {

	var t token.Tokens

	if err := ts.db.GetDb().First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return t, nil
		} else {
			return t, fmt.Errorf("Error al obtener token: %w", err)
		}
	}

	return t, nil
}

func (ts *TokenService) DeleteToken() {
	result := ts.db.GetDb().Exec("DELETE FROM tokens")

	if result.Error != nil {
		fmt.Println("Error eliminando el token:", result.Error)
	} else {
		fmt.Println("Token eliminado exitosamente")
	}
}
