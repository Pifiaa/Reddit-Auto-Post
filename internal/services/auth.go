package services

import (
	request "RedditAutoPost/internal/http"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(url string, c *gin.Context) {
	var firstCredential, err = GetCredentials()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}

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

	request.Post(url, headers, data, c)
}
