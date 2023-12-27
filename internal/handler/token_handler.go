package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database/models/token"
	"RedditAutoPost/internal/request"
	"RedditAutoPost/internal/services"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context, cfg *config.Config) (string, error) {
	tokenService, err := services.NewTokenService()
	if err != nil {
		return "", fmt.Errorf("error al crear el servicio de tokens: %v", err)
	}

	token, err := tokenService.GetToken()
	if err != nil {
		return "", fmt.Errorf("error al obtener el token: %v", err)
	}

	if !tokenIsValid(token) {
		return requestByToken(c, cfg)
	}

	// c.JSON(200, token.Token)

	return token.Token, nil
}

func secondsToDate(seconds float64) time.Time {
	return time.Now().Add(time.Duration(int64(seconds)) * time.Second)
}

func tokenIsValid(tok token.Tokens) bool {
	return time.Now().Unix() < tok.Expiration.Unix()
}

func requestByToken(c *gin.Context, cfg *config.Config) (string, error) {
	tokenService, err := services.NewTokenService()
	if err != nil {
		return "", fmt.Errorf("error al crear el servicio de tokens: %v", err)
	}

	redditCredential, err := services.GetCredentials()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/access_token", cfg.Reddit.Url)
	authString := base64.StdEncoding.EncodeToString([]byte(redditCredential.ClientID + ":" + redditCredential.ClientSecret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s",
		redditCredential.Username,
		redditCredential.Password,
	)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	status, result := request.Post(url, headers, data, c)

	c.JSON(status, result)

	if status == 200 {
		accessToken, ok := result["access_token"].(string)
		if !ok {
			log.Println("No se encontró el token de acceso")
			return "", nil
		}

		date := secondsToDate(result["expires_in"].(float64))
		tokenService.CreateAccessToken(accessToken, date)

		return accessToken, nil
	}

	log.Printf("La solicitud no fue exitosa. Estado: %d\n", status)
	return "", nil
}
