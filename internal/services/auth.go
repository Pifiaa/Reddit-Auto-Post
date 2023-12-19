package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(url string, c *gin.Context) {
	var redditCredential, err = GetCredentials()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err})
	}

	if len(redditCredential) == 0 {
		err = fmt.Errorf("No hay credenciales de acceso disponibles")
		c.JSON(http.StatusNotFound, gin.H{"Error: ": err})
	}
	c.JSON(http.StatusAccepted, redditCredential)

	/*username := redditCredential[0].Username
	password := redditCredential[0].Password
	client_secret := redditCredential[0].ClientSecret
	client_id := redditCredential[0].ClientID
	authString := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))

	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + authString,
	}

	status, result := request.Post(url, headers, data, c)

	if status == 200 {
		c.JSON(status, result)
	} /*else {
		c.JSON(status, result)
	}*/
}
