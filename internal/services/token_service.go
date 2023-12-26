package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/database/models/token"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateAccessToken(accessToken string, expiration time.Time) {
	connection, err := database.DatabaseConnect()
	if err != nil {
		fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	defer connection.Close()

	db := connection.GetDb()

	newToken := token.Tokens{Token: accessToken, Expiration: expiration}

	result := db.Where(token.Tokens{Token: accessToken}).Assign(&newToken).FirstOrCreate(&newToken)

	if result.Error != nil {
		panic("Failed to perform firstOrCreate: " + result.Error.Error())
	}

	fmt.Println("Resullt", result)
}

func GetToken() (token.Tokens, bool) {
	connection, err := database.DatabaseConnect()
	if err != nil {
		return token.Tokens{}, false
	}
	defer connection.Close()

	db := connection.GetDb()

	var token token.Tokens

	if err := db.Unscoped().First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return token, false
		}
		return token, false
	}

	return token, true
}
