package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/models/credentials"
	"fmt"
)

func GetCredentials() (credentials.Credentials, error) {
	db, err := database.Connect()

	if err != nil {
		fmt.Errorf("Error al conectar la base de datos: %w", err)
		return credentials.Credentials{}, err
	}

	var redditCredential []credentials.Credentials
	db.Unscoped().Find(&redditCredential)

	if len(redditCredential) == 0 {
		fmt.Errorf("No hay credenciales de acceso disponibles")
		return credentials.Credentials{}, err
	}

	return redditCredential[0], nil
}
