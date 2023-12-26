package handler

import (
	"RedditAutoPost/internal/database/models/credentials"
	"RedditAutoPost/internal/services"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetCredentials(c *gin.Context) {
	redditCredential, err := services.GetCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}

	c.JSON(http.StatusOK, redditCredential)
}

func CreateCredentials(c *gin.Context) {
	username := "Piifia"
	password := "123456789asd*"
	clientId := "9vDiuDZcj1CNfWBn7hhUbw"
	clientSecret := "1ptla1RhgWKykD9f53RaSyfVHj-FgA"

	credentialFields := credentials.Credentials{
		Username:     username,
		Password:     password,
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}

	err := validateFields(credentialFields)
	if err != nil {
		fmt.Println(err)
	}

	services.CreateCredentials(credentialFields)
}

func validateFields(credentials interface{}) error {
	fields := reflect.ValueOf(credentials)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		value := fields.Field(i).Interface()

		if value == "" {
			return fmt.Errorf("El campo '%s' no puede estar vacio", field.Name)
		}
	}

	return nil
}
