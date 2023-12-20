package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/models/credentials"
	"RedditAutoPost/internal/models/token"
	"fmt"
	"time"
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

func CreateCredentials(fields credentials.Credentials) {
	db, err := database.Connect()

	if err != nil {
		err = fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	// credential := credentials.Credentials{ClientID: fields.ClientID}

	result := db.FirstOrCreate(&fields, fields)

	if result.Error != nil {
		panic("Failed to create or retrieve user")
	}
}

func CreateAccessToken(accessToken string, expiration time.Time) {
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
