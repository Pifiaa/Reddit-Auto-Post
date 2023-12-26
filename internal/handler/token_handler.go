package handler

import (
	"RedditAutoPost/config"
	"RedditAutoPost/internal/database/models/token"
	"RedditAutoPost/internal/request"
	"RedditAutoPost/internal/services"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context, cfg *config.Config) (string, error) {

	token, existToken := services.GetToken()

	if !existToken {
		return requestByToken(c, cfg)
	}

	isValid := tokenExpiration(token)

	if !isValid {
		return requestByToken(c, cfg)
	}

	c.JSON(200, token.Token)

	return token.Token, nil
}

func secondsToDate(seconds float64) time.Time {
	currentTime := time.Now()
	newTime := currentTime.Add(time.Duration(int64(seconds)) * time.Second)

	return newTime
}

func tokenExpiration(token token.Tokens) bool {
	currentTime := time.Now()
	expiration := token.Expiration

	if currentTime.Unix() < expiration.Unix() {
		return true
	}

	return false
}

func requestByToken(c *gin.Context, cfg *config.Config) (string, error) {
	redditCredential, err := services.GetCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return "", nil
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

		accessToken, accessTokenExists := result["access_token"].(string)

		if accessTokenExists {
			date := secondsToDate(result["expires_in"].(float64))
			services.CreateAccessToken(accessToken, date)

			return accessToken, nil
		} else {
			fmt.Println("No se encontró el token de acceso en la respuesta.")
			return "", nil
		}

	} else {
		fmt.Println("La solicitud no fue exitosa. Estado:", status)
		return "", nil
	}
}
