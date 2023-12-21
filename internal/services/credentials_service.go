package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/models/credentials"
	"RedditAutoPost/internal/models/token"
	"fmt"
)

func GetCredentials() ([]credentials.Credentials, error) {
	db, err := database.Connect()

	if err != nil {
		err = fmt.Errorf("Error al conectar la base de datos: %w", err)
		return []credentials.Credentials{}, err
	}

	var redditCredential []credentials.Credentials
	db.Limit(1).Unscoped().Find(&redditCredential)

	database.Close()

	return redditCredential, nil
}

func CreateCredentials(fields credentials.Credentials) {
	db, err := database.Connect()

	if err != nil {
		err = fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	// Utiliza FirstOrCreate con un puntero a la instancia de Credentials
	result := db.Unscoped().FirstOrCreate(&fields, credentials.Credentials{ClientID: fields.ClientID})

	if result.Error != nil {
		// Manejar el error de alguna manera
		fmt.Println("Error al crear o buscar credenciales:", result.Error)
		return
	}

	// Imprime el resultado
	fmt.Println("Credenciales agregadas")
}

func CreateAccessToken(accessToken string, expiration string) {
	db, err := database.Connect()

	if err != nil {
		err = fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	newToken := token.AccessToken{Token: accessToken, Expiration: expiration}

	result := db.Where(token.AccessToken{Token: accessToken}).Assign(&newToken).FirstOrCreate(&newToken)

	if result.Error != nil {
		panic("Failed to perform firstOrCreate: " + result.Error.Error())
	}

	fmt.Println("Resullt", result)
}
