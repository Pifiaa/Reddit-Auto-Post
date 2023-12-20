package handler

import (
	"RedditAutoPost/config"
	request "RedditAutoPost/internal/http"
	"RedditAutoPost/internal/services"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	redditCredential, err := services.GetCredentials()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if len(redditCredential) == 0 {
		err = fmt.Errorf("No hay credenciales de acceso disponibles")
		c.JSON(http.StatusNotFound, gin.H{"Error: ": err.Error()})
		return
	}

	username := redditCredential[0].Username
	password := redditCredential[0].Password
	client_secret := redditCredential[0].ClientSecret
	client_id := redditCredential[0].ClientID
	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	url := fmt.Sprintf("%s/access_token", config.GetEnv("reddit.url"))

	status, result := request.Post(url, headers, data, c)

	if status == 200 {
		expirationSeconds, _ := strconv.Atoi(result["expires_in"].(string))
		expiration := time.Now().Add(time.Second * time.Duration(expirationSeconds))

		services.CreateAccessToken(result["access_token"].(string), expiration)
	}
}
