package services

import (
	"RedditAutoPost/internal/database"
	"RedditAutoPost/internal/http"
	"RedditAutoPost/internal/models/credentials"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(url string, c *gin.Context) {
	db, err := database.Connect()

	if err != nil {
		fmt.Errorf("Error al conectar la base de datos: %w", err)
	}

	var redditCredential []credentials.Credentials
	db.Unscoped().Find(&redditCredential)

	if len(redditCredential) == 0 {
		fmt.Errorf("No hay credenciales de acceso disponibles")
		return
	}

	firstCredential := redditCredential[0]

	username := firstCredential.Username
	password := firstCredential.Password
	client_secret := firstCredential.ClientSecret
	client_id := firstCredential.ClientID

	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	http.Post(url, headers, data, c)
}
