package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/models/credentials"
	"fmt"
)

func GetCredentials() ([]credentials.Credentials, error) {
	db, err := database.Connect()

	if err != nil {
		err = fmt.Errorf("Error al conectar la base de datos: %w", err)
		return []credentials.Credentials{}, err
	}

	var redditCredential []credentials.Credentials
	db.Unscoped().Find(&redditCredential)

	database.Close()

	return redditCredential, nil
}
