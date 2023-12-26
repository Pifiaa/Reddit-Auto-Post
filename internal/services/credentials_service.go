package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/database/models/credentials"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetCredentials() (credentials.Credentials, error) {
	connection, err := database.DatabaseConnect()
	if err != nil {
		return credentials.Credentials{}, fmt.Errorf("Error al conectar la base de datos: %w", err)
	}
	defer connection.Close()

	db := connection.GetDb()

	var redditCredential credentials.Credentials

	if err := db.Unscoped().First(&redditCredential).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return credentials.Credentials{}, fmt.Errorf("No se encontraron credenciales: %w", err)
		}
		return credentials.Credentials{}, fmt.Errorf("Error al obtener credenciales: %w", err)
	}

	return redditCredential, nil
}

func CreateCredentials(fields credentials.Credentials) {
	connection, err := database.DatabaseConnect()

	if err != nil {
		fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	defer connection.Close()

	db := connection.GetDb()

	result := db.Unscoped().FirstOrCreate(&fields, credentials.Credentials{ClientID: fields.ClientID})

	if result.Error != nil {
		// Manejar el error de alguna manera
		fmt.Println("Error al crear o buscar credenciales:", result.Error)
		return
	}

	// Imprime el resultado
	fmt.Println("Credenciales agregadas")
}
